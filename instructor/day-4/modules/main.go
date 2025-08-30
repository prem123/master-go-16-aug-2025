package main

import (
	"fmt"

	"github.com/the-code-camp/master-go-16-aug-2025/utils"
	"rsc.io/quote"
)

func main() {
	result := utils.Sum(10, 20)
	fmt.Println(result)

	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
}
