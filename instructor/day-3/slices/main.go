package main

import "fmt"

func main() {
	// var nums []int // slice does not have the size defined
	nums := make([]int, 1)

	fmt.Println("Len: ", len(nums))
	fmt.Println("Values: ", nums)

	nums[0] = 10
	nums = append(nums, 100, 101, 102)

	fmt.Println("Len: ", len(nums))
	fmt.Println("Values: ", nums)

	// stepping...
	for i := 0; i < len(nums); i += 3 {
		fmt.Println(nums[i])
	}
}
