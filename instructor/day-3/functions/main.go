package main

import "fmt"

func main() {
	// result := sum(1, 2)

	n, err := fmt.Println("hello")

	// Println supports varargs argument represented by ...
	fmt.Println("Bytes and error: ", n, err)

}

func print(msg string) {
	s := fmt.Sprintf("Hello, the message is %s", msg)
	fmt.Println(s)
}

// (paramaters) return
func sum(a int, b int) (result int, n int) {
	// var result int
	result = a + b
	return result, 0
}

// func anotherSum(a, b, c int, msg, msg2 string) int {
// 	return a + b
// }
