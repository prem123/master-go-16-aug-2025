package main

import "fmt"

func main() {
	// p := "New Delhi"
	// p := 10
	p := 10.5
	fmt.Printf("%T\n", p)

	x, y := 65, 98

	fmt.Printf("%c %c \n", x, y)

	msg := "some string value"
	fmt.Printf("%s\n", msg)

	a, b, c := 10.0, 2012.23, 321.9

	// Applying padding, same can be done with string and other types
	fmt.Printf("%-7.2f %-7.2f\n", a, b)
	fmt.Printf("%-7.2f\n", b)
	fmt.Printf("%-7.2f\n", c)

	// Sprintf
	// s := "Hello <<name>>, your age is <<x>> years"
	s := fmt.Sprintf("Hello %s, your age is %d years", "Rishi", 30)
	fmt.Println(s)

}
