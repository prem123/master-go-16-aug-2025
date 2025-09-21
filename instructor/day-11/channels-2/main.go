package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() { // schedule for future
		ch <- 10 // SEND
	}()

	data := <-ch // RECEIVE

	fmt.Println("Result: ", data)
}
