# marathon

[![Go](https://github.com/dotWicho/marathon/workflows/Go/badge.svg?branch=master)](https://github.com/dotWicho/marathon)
[![Quality Report](https://goreportcard.com/badge/github.com/dotWicho/marathon)](https://goreportcard.com/badge/github.com/dotWicho/marathon)
[![GoDoc](https://godoc.org/github.com/dotWicho/marathon?status.svg)](https://pkg.go.dev/github.com/dotWicho/marathon?tab=doc)

## Library to manage Marathon servers via API Calls

## Getting started

- API documentation is available via [godoc](https://godoc.org/github.com/dotWicho/marathon).
- Test code contains some small examples of the use of this library.

## Installation

To install Marathon package, you need to install Go and set your Go workspace first.

1 - The first need [Go](https://golang.org/) installed (**version 1.13+ is required**).
Then you can use the below Go command to install Marathon

```bash
$ go get -u github.com/dotWicho/marathon
```

And then Import it in your code:

``` go
package main

import "github.com/dotWicho/marathon"
```
Or

2 - Use as module in you project (go.mod file):

``` go
module myclient

go 1.13

require (
	github.com/dotWicho/marathon v1.4.4
)
```

## Contributing

- Get started by checking our [contribution guidelines](https://github.com/dotWicho/marathon/blob/master/CONTRIBUTING.md).
- Read the [dotWicho marathon wiki](https://github.com/dotWicho/marathon/wiki) for more technical and design details.
- If you have any questions, just ask!

