package main

import "fmt"

func main() {
	// rune
	// for ch := 'a'; ch <= 'z'; ch++ {
	// 	fmt.Printf("%d ", ch)
	// }

	// for <index>, <val> := range <string, array, slice>

	// for _, val := range "Hello" {
	// 	fmt.Printf("%c\n", val)
	// }

	for row := 0; row < 5; row++ {
		for col := 0; col <= row; col++ {
			fmt.Print("* ") // it will print, you cannot format and there is no new line at the end.
			// fmt.Printf("* ") // formatting
		}
		fmt.Println() // at the end prints with new line
	}
}
