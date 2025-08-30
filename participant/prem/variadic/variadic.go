package main

import "fmt"

// max function that accepts zero, one, or many integers using variadic parameters
func max(nums ...int) int {
	// If no numbers are passed, return 0
	if len(nums) == 0 {
		fmt.Println("No numbers provided, returning 0")
		return 0
	}

	// Start with the first number as the maximum
	maximum := nums[0]

	// Compare with all other numbers
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
	}

	return maximum
}

func main() {
	fmt.Println("=== Variadic Function Demo ===\n")

	// Case 1: No numbers
	fmt.Println("Case 1: Calling max() with no numbers")
	result1 := max()
	fmt.Printf("Result: %d\n\n", result1)

	// Case 2: Three numbers
	fmt.Println("Case 2: Calling max() with three numbers (10, 25, 15)")
	result2 := max(10, 25, 15)
	fmt.Printf("Result: %d\n\n", result2)

	// Case 3: A slice expanded with ...
	fmt.Println("Case 3: Calling max() with a slice expanded using ...")
	numbers := []int{5, 42, 18, 7, 33, 12}
	fmt.Printf("Slice: %v\n", numbers)
	result3 := max(numbers...)
	fmt.Printf("Result: %d\n\n", result3)

	// Additional examples
	fmt.Println("Additional examples:")

	// Single number
	fmt.Printf("max(100): %d\n", max(100))

	// Many numbers
	fmt.Printf("max(1, 2, 3, 4, 5, 6, 7, 8, 9): %d\n", max(1, 2, 3, 4, 5, 6, 7, 8, 9))

	// Negative numbers
	fmt.Printf("max(-10, -5, -20, -1): %d\n", max(-10, -5, -20, -1))
}

