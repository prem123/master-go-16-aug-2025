## Flags

The `flag` package in Go (Golang) is a built-in standard library package used for parsing command-line arguments, commonly referred to as "flags," in Go programs. It provides a straightforward way to define, parse, and access values passed to your application during runtime.


```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define a string flag with a default value and a help message
	name := flag.String("name", "World", "a name to say hello to")

	// Define an integer flag
	age := flag.Int("age", 30, "the age of the person")

	// Parse the command-line arguments
	flag.Parse()

	// Access and use the flag values
	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)

	// Access positional arguments
	if len(flag.Args()) > 0 {
		fmt.Printf("Positional arguments: %v\n", flag.Args())
	}
}
```