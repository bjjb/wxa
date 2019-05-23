# Λ

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

import Λ "github.com/bjjb/wxa"

func main() {
  λ := func(i interface{}) interface{} {
    return i.(int) * 2
  }
  fmt.Println(Λ.Map(λ, 1, 2, 3, 4, 5))
  // Output:
  // 2 4 6 8 10

  multer := func(acc, next interface{}) interface{} {
    return acc.(int) * acc.(next)
  }
  fmt.Println(Λ.Reduce(λ, 1, 2, 3, 4, 5))
  // Output:
  // 120

  λ := func(v interface{}) bool {
    switch t := v.(type) {
    case bool:
      return v.(bool)
    case int:
      return v.(int) != 0
    case string:
      return len(v.(string)) > 0
    case nil:
      return false
    }
  }

  fmt.Println(Λ.Any(truthy, 0, "", []string{}, false))
  // Output:
  // false
}
```

Documentation is available on [GoDoc](https://godoc.org/gitlab.com/bjjb/λ).

[home]: https://travis-ci.com/bjjb/wxa
[badge]: https://travis-ci.com/bjjb/wxa.svg?branch=master
