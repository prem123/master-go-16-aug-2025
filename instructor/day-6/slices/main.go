package main

import (
	"fmt"
	"slices"
)

// TODO: difference between range and standard for loop

func main() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println(nums, len(nums), cap(nums))

	n1 := nums[2:4] //re-slicing
	// n1 and nums are pointing to the same underlying array at different index positions
	fmt.Println(n1, "len: ", len(n1), "cap: ", cap(n1))
	n1 = append(n1, 99)

	fmt.Println("n1: ", n1)
	fmt.Println("nums: ", nums)

	fmt.Println("n1-> ", "len: ", len(n1), "cap: ", cap(n1))
	n1 = append(n1, 88)
	fmt.Println("n1-> ", n1, "len: ", len(n1), "cap: ", cap(n1))
	n1 = append(n1, 89)
	fmt.Println("n1-> ", n1, "len: ", len(n1), "cap: ", cap(n1))
	n1 = append(n1, 90)
	fmt.Println("n1-> ", n1, "len: ", len(n1), "cap: ", cap(n1))
	n1 = append(n1, 91)
	fmt.Println("n1-> ", n1, "len: ", len(n1), "cap: ", cap(n1))
	//Delete(slice, starting-index, ending-index)
	nums = slices.Delete(nums, 2, 4)
	fmt.Println(nums)
	// fmt.Println("nums-> ", nums, "len: ", len(nums), "cap: ", cap(nums))
	// multiplyByTwoSlice(nums)

	// fmt.Println(nums)

}

// does not work because Go is pass by value
func multiplyByTwo(nums [5]int) {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * 2
	}
	// for _, n := range nums {
	// 	n = n * 2
	// }
}

// whenever mutating slices, always use indexes
// Go is pass by value in most of the cases. To understand this we need to
// understand slices better.
func multiplyByTwoSlice(nums []int) {
	for i, _ := range nums {
		nums[i] = nums[i] * 2
	}
}
