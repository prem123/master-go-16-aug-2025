package main

import "fmt"

func main() {
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("===========")
	// i = 5
	// starting at this point value of i is 5
	// while loop
	for i > 0 {
		fmt.Println(i)
		i--
	}

	fmt.Println("===========")
	// infinite loop - just like while true in other lang
	for {
		fmt.Println(i)
		if i == 5 {
			break
		}
		i++
	}
}
