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

Below is a complete annotated example to generate multiple keys given a path template.

```go
package main

import (
        "fmt"
        "os"

        hdpath "github.com/mcdee/go-hdpath"
)

func main() {
        // The template uses 'n' to signify the instance; other values are static.
        template, err := hdpath.Parse("m/44'/0'/n'/0/0")
        if err != nil {
                fmt.Fprintf(os.Stderr, "Failed to parse template: %v\n", err)
        }

        for i := range uint32(5) {
                path := template.Instance(i)
                fmt.Fprintf(os.Stdout, "Instance %d path is %s with values %v\n", i, path.String(), path.Values())
        }
}
```

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/mcdee/go-hdpath/issues).

## License

[Apache-2.0](LICENSE) © 2025 Jim McDonald
