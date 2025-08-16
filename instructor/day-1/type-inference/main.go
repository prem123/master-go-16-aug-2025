package main

import "fmt"

var x int = 10

func main() {
	fmt.Println(x)

	var i = 10
	i = 20

	var b = "abc"

	fmt.Println(i, b)

	var msg string
	fmt.Println(msg + " World")

	// int8 cannot store value larger than 127
	var x int8 = 127
	// uint cannot store a negative numbers (signed)
	var y uint = 7
	println(x, y)

	// infering type and creating variable
	j := 20
	message := "Welcome to type inference"

	// message := "Welcome" // cannot do this again because the variable is already defined

	fmt.Println(j, message)
}
