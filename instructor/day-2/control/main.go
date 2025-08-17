package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // another syntax with initialization
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switch_function() {
	// continue_func()

	var i = 3
	switch i {
	case 1:
		println("value is one")
	case 2:
		println("value is two")
	default:
		println("default case")
	}
}

func continue_func() {
	i := 1 // this syntax is more favoured

	for ; i < 11; i++ {
		// If statements in Go are often used as guard clauses.
		if i%2 == 1 {
			continue
		}
	}
	fmt.Println(i)
}
