package main

import (
	"fmt"
	"io"

	"github.com/cli/go-gh/v2/pkg/config"
)

// Config is the configuration for this extension.
type Config struct {
	Source string
}

var defaultConfig = &Config{
	Source: "spenserblack/.gitattributes",
}

type ghConfigReader = func() (*config.Config, error)

func newConfig(r ghConfigReader) (*Config, error) {
	cfg := &Config{}
	ghConfig, err := r()
	if err != nil {
		return nil, err
	}
	// NOTE This isn't nested, but nesting is awkward from the CLI with `gh config set`
	source, err := ghConfig.Get([]string{"gh_gitattributes_source"})
	if _, ok := err.(*config.KeyNotFoundError); ok {
		source = defaultConfig.Source
		err = nil
	} else if err != nil {
		return nil, err
	}
	cfg.Source = source
	return cfg, nil
}

func newConfigOrDefault(stderr io.Writer) *Config {
	cfg, err := newConfig(config.Read)
	if err != nil {
		fmt.Fprintf(stderr, "error reading config: %s\nUsing default config", err)
		return defaultConfig
	}
	return cfg
}
