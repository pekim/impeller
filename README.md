# impeller

[![Go Reference](https://pkg.go.dev/badge/github.com/pekim/impeller.svg)](https://pkg.go.dev/github.com/pekim/impeller)
[![golangci-lint](https://github.com/pekim/impeller/actions/workflows/verify.yml/badge.svg)](https://github.com/pekim/impeller/actions/workflows/verify.yml)

This library provides cgo-free Go bindings for
[Impeller](https://github.com/flutter/flutter/tree/master/engine/src/flutter/impeller).

## status

This library is experimental.
It appears to broadly work (at least on linux),
but it has not been extensively tested.

## AI

No AI was used in the creation of this library.

## development

### upgrading Impeller

- update the `FLUTTER_SHA` env var in
  `internal/update.sh` to the desired
  [flutter repo](https://github.com/flutter/flutter)
  git commit
- run `./internal/update`, to download and extract
  the artifacts
- run `go generate` to regenerate the Go binding
- once satisfied that all is good, commit

### pre-commit hook

There are configuration files for linting and other checks.
To use a git pre-commit hook for the checks

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` (v2.x) if not already installed
  - https://golangci-lint.run/docs/welcome/install/#binaries
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
