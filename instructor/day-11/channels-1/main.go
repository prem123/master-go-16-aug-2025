package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go add(10, 20, wg, ch)

	wg.Wait()
	result := <-ch // RECEIVE
	fmt.Println("Result: ", result)
}

func add(a, b int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- a + b // SEND it to channel
}
