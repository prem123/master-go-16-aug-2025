package main

import (
	"fmt"
)

func main() {

	var input int
	var x, y float64
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Multiply")
		fmt.Println("3. Exit")

		fmt.Printf("Enter your choice:")
		fmt.Scanln(&input)

		if input < 1 || input > 3 {
			continue
		} else if input == 3 {
			break
		}

		fmt.Printf("Enter first operand:")
		fmt.Scanln(&x)
		fmt.Printf("Enter second operand:")
		fmt.Scanln(&y)

		switch input {
		case 1:
			fmt.Println("Sum is: ", x+y)
		case 2:
			fmt.Println("Multiply is: ", x*y)
		default:
			fmt.Println("invalid choice...")
		}

	}
}
