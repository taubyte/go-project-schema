# go-project-schema

[![Release](https://img.shields.io/github/release/taubyte/go-project-schema.svg)](https://github.com/taubyte/go-project-schema/releases)
[![License](https://img.shields.io/github/license/taubyte/go-project-schema)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/taubyte/go-project-schema)](https://goreportcard.com/report/taubyte/go-project-schema)
[![GoDoc](https://godoc.org/github.com/taubyte/go-project-schema?status.svg)](https://pkg.go.dev/github.com/taubyte/go-project-schema)
[![Discord](https://img.shields.io/discord/973677117722202152?color=%235865f2&label=discord)](https://tau.link/discord)

This is a golang interface to the Taubyte project using [go-seer](https://github.com/taubyte/go-seer). 

## Testing

[test project](internal/test/config/)

### Run all

```shell
go test -v ./...
```

### Calculate coverage for all packages
```shell
go test -v ./... -coverprofile cover.out -coverpkg ./...
```

#### Display coverage
```shell
# html
go tool cover -html=cover.out

# each function % and total %
go tool cover -func=cover.out
```

# Maintainers
 - Samy Fodil @samyfodil
 - Sam Stoltenberg @skelouse
 - Tafseer Khan @tafseer-khan
 - Aron Jalbuena @arontaubyte
