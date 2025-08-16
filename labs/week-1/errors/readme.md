## Errors in Go

Errors are an essential part of programming, and Go handles them differently compared to many other languages. Instead of exceptions, Go uses explicit error values returned from functions. This makes error handling predictable and explicit.

1. Understanding Errors in Go

The `error` type

In Go, errors are represented by the built-in `error` type.

It is simply an `interface` with one method:

```go
type error interface {
    Error() string
}
```

Example: Returning an error

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

2. Creating Custom Errors
You can create your own error messages:

```go
err := errors.New("something went wrong")
```
Or format them dynamically:
```go
fmt.Errorf("failed to open file: %s", filename)
```


### Lab: Error (do this lab after function)

Create a function `validateAge(age int)`

If `age < 18`, return a validation error otherwise return `nil`.

Test with different ages.