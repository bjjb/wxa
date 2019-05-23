# λ

[![Build Status][badge]][home]

A set of potentially useful higher-order functions in Go. Write some functions
which use integers, and use them with `Λ.Map`, `Λ.Reduce`, `Λ.Find`, etc.

## Installation

    go get -u github.com/bjjb/wxa

## Usage

```go
import Λ "gitlab.com/bjjb/wxa"
```

## Examples

```go
package main

import (
  "fmt"
  "github.com/bjjb/wxa"
)

func main() {
  λ := func(i interface{}) (interface{}, error) {
    return i.(int) * 2, nil
  }

  f := wxa.Map(doubler)

  for i, r := range f(1, 2, 3, 4, 5) {
    fmt.Println("%d: %d", i, r)
  }
}
```

Documentation is available on [GoDoc](https://godoc.org/github.com/bjjb/wxa).

[home]: https://travis-ci.com/bjjb/wxa
[badge]: https://travis-ci.com/bjjb/wxa.svg?branch=master
