package main

import (
	"fmt"
	"time"
)

func main() {

	defer fmt.Println("Good bye!!")

	input := 0
	fmt.Print("Enter the number in seconds: ")
	fmt.Scanln(&input)

	if input <= 0 {
		return // break works only inside loop
	}
	for ; input > 0; input-- {
		// for i := input; i > 0; i-- {
		fmt.Println(input)
		time.Sleep(1 * time.Second)
	}
}
