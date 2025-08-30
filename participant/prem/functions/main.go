package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Variadic Functions Lab ===")
	testVariadicFunctions()
	
	fmt.Println("\n=== Interactive Calculator (Refactored) ===")
	runCalculator()
}

func testVariadicFunctions() {
	fmt.Printf("max() = %d\n", max())
	fmt.Printf("max(5, 2, 8, 1) = %d\n", max(5, 2, 8, 1))
	
	numbers := []int{10, 25, 3, 99, 42}
	fmt.Printf("max(numbers...) = %d\n", max(numbers...))
	
	fmt.Printf("concat() = '%s'\n", concat())
	fmt.Printf("concat(\"Hello\", \"World\") = '%s'\n", concat("Hello", "World"))
	fmt.Printf("concat(\"Go\", \"is\", \"awesome\") = '%s'\n", concat("Go", "is", "awesome"))
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	
	maximum := nums[0]
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

func concat(words ...string) string {
	return strings.Join(words, " ")
}

func runCalculator() {
	for {
		choice := displayMenuAndGetChoice()
		
		if choice == 5 {
			fmt.Println("Goodbye!")
			break
		}
		
		if !isValidChoice(choice) {
			fmt.Println("Invalid choice")
			continue
		}
		
		operand1, operand2 := getOperands()
		result := performOperation(choice, operand1, operand2)
		displayResult(choice, operand1, operand2, result)
	}
}

func displayMenuAndGetChoice() int {
	fmt.Println("\n--- Calculator Menu ---")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func isValidChoice(choice int) bool {
	return choice >= 1 && choice <= 5
}

func getOperands() (float64, float64) {
	var operand1, operand2 float64
	
	fmt.Print("Enter first operand: ")
	fmt.Scanln(&operand1)
	
	fmt.Print("Enter second operand: ")
	fmt.Scanln(&operand2)
	
	return operand1, operand2
}

func performOperation(choice int, operand1, operand2 float64) float64 {
	switch choice {
	case 1:
		return add(operand1, operand2)
	case 2:
		return subtract(operand1, operand2)
	case 3:
		return multiply(operand1, operand2)
	case 4:
		return divide(operand1, operand2)
	default:
		return 0
	}
}

func displayResult(choice int, operand1, operand2, result float64) {
	operations := map[int]string{
		1: "Addition",
		2: "Subtraction", 
		3: "Multiplication",
		4: "Division",
	}
	
	symbols := map[int]string{
		1: "+",
		2: "-",
		3: "*",
		4: "/",
	}
	
	fmt.Printf("%s: %.2f %s %.2f = %.2f\n", 
		operations[choice], operand1, symbols[choice], operand2, result)
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Division by zero!")
		return 0
	}
	return a / b
}
