# `gh gitattributes`

[![CI](https://github.com/spenserblack/gh-gitattributes/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/gh-gitattributes/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/gh-gitattributes/branch/main/graph/badge.svg?token=xYXvCen5Un)](https://codecov.io/gh/spenserblack/gh-gitattributes)

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

- `gh_gitattributes_source`: The default source repository to use

## Installation

```shell
gh extension install spenserblack/gh-gitattributes
```
