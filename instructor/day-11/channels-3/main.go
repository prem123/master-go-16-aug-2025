package main

import "fmt"

func main() {
	ch := make(chan int)

	// THIS WILL RESULT IN DEADLOCK
	data := <-ch // RECEIVE

	go func() { // schedule for future
		ch <- 10 // SEND
	}()

	fmt.Println("Result: ", data)
}
