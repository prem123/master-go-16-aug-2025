package main

import "fmt"

func main() {
	// anonymous functions in Go
	// you have to consume the anonymous function immediately
	// mostly used as go routines

	func() {
		fmt.Println("inside anonymous function...")
	}()

	// in here sum is a function
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(10, 20))

	// In below example, result will the total of 10 and 20 and result is int
	result := func(a, b int) int {
		return a + b
	}(20, 20)
	fmt.Println("Result: ", result)

}
