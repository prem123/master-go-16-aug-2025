package main

import "fmt"

func main() {

	brothers := []string{"chico", "harpo", "groucho", "gummo", "zeppo"}
	fmt.Println(brothers)

	wellknown := brothers[1:3]
	fmt.Println(wellknown)

	// lastTwo := brothers[3:5] // works well
	lastTwo := brothers[3:]
	fmt.Println(lastTwo)

	firstTwo := brothers[:3]
	fmt.Println(firstTwo)

	s := "hello" // strings are sequence of rune

	fmt.Println(s[:3])
}
