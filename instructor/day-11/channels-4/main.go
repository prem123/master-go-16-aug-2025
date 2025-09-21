package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() { // schedule for future
		defer wg.Done()
		data := <-ch // RECEIVE
		fmt.Println("Result: ", data)
	}()

	ch <- 10 // SEND
	wg.Wait()
}
