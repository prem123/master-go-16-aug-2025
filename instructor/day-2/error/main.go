package main

import (
	"errors"
	"fmt"
)

func main() {
	// val, err := strconv.Atoi("abc")

	// fmt.Println(val)

	a := 10.0
	b := 0.0
	result, err := divide(a, b)

	if err != nil {
		// that means there is an error
		// take decision, stop program or return error to the caller
		fmt.Println("zero div error: ", err)
	} else {
		// otherwise all is fine go ahead.
		fmt.Println(result)
	}

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("will result in div by zero panic")
	} else {
		return a / b, nil
	}
}
