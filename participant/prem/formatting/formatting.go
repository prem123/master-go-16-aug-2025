package main

import "fmt"

func main() {
	p := "New Delhi"
	fmt.Printf("%T\n", p)
	charA, charZ := 97, 122
	fmt.Printf("type: %c and %c\n", charA, charZ)
	fmt.Printf("srt1: %s\n", "\"string\"")
	fmt.Printf("str2: %q\n", "\"string\"")
	fmt.Printf("width1: |%6d|%6d|\n", 19, 731)
	fmt.Printf("width2: |%.2f|%.2f|\n", 6.7, 12.82)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.0)

	chr := 97
	str := fmt.Sprintf("%c", chr)
	fmt.Println(str)
	s := fmt.Sprintf("\"sprintf: a %s\"")
	fmt.Println(s)
}
