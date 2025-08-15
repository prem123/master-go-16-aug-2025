## Zero value
In previous examples we've written code like this

```go
var name = "go"
var counter = 1
```

Which both declares and initalises the variables counter and name respectively. 

📋 What happens if we have code like this?

```go
package main

import "fmt"

func main() {
  var name string
  var counter int
  fmt.Printf("%q, %v", name, counter)
}
```

In Go, there is no unitialised memory. The Go runtime will always ensure that the memory allocated for each variable is initalised before use. If we write something like:

```go
var name string
var counter int
```

Then the memory assigned to the variables name and counter will be zeroed, as we have not provided an initaliser.
- The value of name will be "" because that is the value of a string with zero length.
- The value of counter will be zero, because that is the value of an `int` if we wrote 0 to its memory location.

Every type in Go has an associated zero value. The value of that variable if we wrote zeros to its memory.
- The zero value for integer types: `int`, `int8`, `uint`, `uint64`, etc, is 0.
- The zero value for floating point types: `float32`, `float64`, `complex128`, etc, is 0.0.
- The zero value for `arrays` is the zero value for each element, ie. `[3]int` is 0, 0, 0.
- The zero value for `slices` is nil.
- The zero value for `structs` is the zero value for each field.