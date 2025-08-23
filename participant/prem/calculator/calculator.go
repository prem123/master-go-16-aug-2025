package main

import "fmt"

func main() {
	for {

		fmt.Println("\n=== Calculator ===")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		var input int
		fmt.Printf("Enter your choice: ")
		fmt.Scanln(&input)

		if input == 5 {
			fmt.Println("Thanks.Exiting")
			break
		}

		if input < 1 || input > 5 {
			fmt.Println("Invalid choice")
			continue
		}

		var operand1, operand2 float64
		fmt.Printf("Enter first operand: ")
		fmt.Scanln(&operand1)
		fmt.Printf("Enter second operand: ")
		fmt.Scanln(&operand2)

		var result float64
		var operation string

		switch input {
		case 1:
			result = operand1 + operand2
			operation = "Addition"
		case 2:
			result = operand1 - operand2
			operation = "Subtraction"
		case 3:
			result = operand1 * operand2
			operation = "Multiplication"
		case 4:
			if operand2 == 0 {
				fmt.Println("Error: Division by zero is not allowed")
				continue
			}
			result = operand1 / operand2
			operation = "Division"
		}

		fmt.Printf("%s result: %.2f\n", operation, result)
	}
}
