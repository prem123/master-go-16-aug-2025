package main

import "fmt"

func main() {
	var nums [5]int // arrays have fixed length

	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50
	// nums[5] = 50 // this is an error, index out of bounds
	fmt.Println("Length: ", len(nums))
	fmt.Println(nums[3])

	nums2 := [5]int{5, 6, 7, 8, 9}
	fmt.Println("nums2: ", nums2)
	fmt.Println("nums2[4]: ", nums2[4])

	sum := 0
	for _, n := range nums2 {
		sum = sum + n
	}
	fmt.Println(sum)
}
