package main

import (
	"fmt"
)

func main() {

	x := 10.0
	y := 20

	fmt.Println(x + float64(y))

	var a float32 = 12.34
	var b float64 = 10.30

	fmt.Println(float64(a) + b) // 22.64000015258789
}
