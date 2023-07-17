package main

import "flag"

var sourceFlag *string
var stdoutFlag *bool

func setupFlags(config *Config) {
	sourceFlag = flag.String("source", config.Source, "The source of the .gitattributes files")
	stdoutFlag = flag.Bool("stdout", false, "Print the .gitattributes file to stdout")
}
