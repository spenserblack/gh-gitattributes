# `gh gitattributes`

[![CI](https://github.com/spenserblack/gh-gitattributes/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/gh-gitattributes/actions/workflows/ci.yml)

Copy `.gitattributes` from an (unofficial) source.

## Description

This pulls a list of files with the `.gitattributes` extension (`FOO.gitattributes`)
and builds a list to prompt you to pick the file you want. If `Common.gitattributes` is
found, you will also be asked if you want to include them.

### CLI

- `-source`: Specify the source repository
- `-stdout`: Write to STDOUT instead of `./.gitattributes`

### Configuration

The configuration values can be set with `gh config set KEY`.

- `gh-gitattributes.source`: The default source repository to use
