package main

import (
	"errors"

	"github.com/cli/go-gh/v2/pkg/config"
)

var invalidGhConfigReader ghConfigReader = func() (*config.Config, error) {
	return nil, errors.New("invalid")
}

func newGhConfigReader(fixture string) ghConfigReader {
	return func() (*config.Config, error) {
		return config.ReadFromString(fixture), nil
	}
}
