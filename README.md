# go-hdpath

[![Tag](https://img.shields.io/github/tag/mcdee/go-hdpath.svg)](https://github.com/mcdee/go-hdpath/releases/)
[![License](https://img.shields.io/github/license/mcdee/go-hdpath.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/mcdee/go-hdpath?status.svg)](https://godoc.org/github.com/mcdee/go-hdpath)
![Lint](https://github.com/mcdee/go-hdpath/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/mcdee/go-hdpath)](https://goreportcard.com/report/github.com/mcdee/go-hdpath)

Go library providing an abstraction for hierarchical deterministic (HD) paths for key generation.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-hdpath` is a standard Go module which can be installed with:

```sh
go get github.com/mcdee/go-hdpath
```

## Usage

Please read the [Go documentation for this library](https://godoc.org/github.com/mcdee/go-hdpath) for interface information.

## Example

Below is a complete annotated example to generate multiple keys given a path.

```go
package main

import (
        "fmt"
        "os"

        hdpath "github.com/mcdee/go-hdpath"
)

func main() {
        // The path string uses 'n' to signify the instance; other values are static.
        path, err := hdpath.Parse("m/44'/0'/n'/0/0")
        if err != nil {
                fmt.Fprintf(os.Stderr, "Failed to parse path: %v\n", err)
        }
        fmt.Fprintf(os.Stdout, "Unresolved path is %s\n", path.String())

        for i := range uint32(5) {
                // Use path.Instance() to create a path with the instance variable resolved.
                resolvedPath := path.Instance(i)
                values, err := resolvedPath.Values()
                if err != nil {
                        // An error will be thrown if an attempt is made to obtain values on a path without a resolved instance.
                        fmt.Fprintf(os.Stderr, "Values() called on an unresolved path: %v\n", err)
                }
                fmt.Fprintf(os.Stdout, "Instance %d path is %s with values %v\n", i, resolvedPath.String(), values)
        }
}
```

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/mcdee/go-hdpath/issues).

## License

[Apache-2.0](LICENSE) © 2025 Jim McDonald
