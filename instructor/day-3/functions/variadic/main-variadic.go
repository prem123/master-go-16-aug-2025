package main

import "fmt"

func main() {

	result := sum()
	fmt.Println(result)

	result = sum(1, 2)
	fmt.Println(result)

	result = sum(1, 2, 4)
	fmt.Println(result)

	result = sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func sum(nums ...int) (result int) {
	for _, n := range nums {
		result = result + n
	}
	return result
}
