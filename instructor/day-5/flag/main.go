package main

import (
	"flag"
	"fmt"
	"time"
)

// usage
// go run main.go --num 1 --name ashish
func main() {

	input := flag.Int("num", 2, "number in seconds")
	name := flag.String("name", "Hello", "your name")

	flag.Parse()

	defer fmt.Printf("%s, Good bye!!\n", *name)

	if *input <= 0 {
		return // break works only inside loop
	}
	for ; *input > 0; *input-- {
		fmt.Println(*input)
		time.Sleep(1 * time.Second)
	}
}
