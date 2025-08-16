package main

import "fmt"

func main() {
	var π = 3.14
	fmt.Println(π)

	var x, y, z = 1, 2, 3

	_ = z
	_ = y

	x++

	fmt.Println(x)
	x--
	fmt.Println(x)

}
