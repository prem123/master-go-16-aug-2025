package main

import (
	"fmt"
	"time"
)

func main() {
	ch := getNos() //scheduled for future

	// consumer
	for data := range ch {
		fmt.Println(data)
	}
}

// producer
func getNos() chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(500 * time.Millisecond) //blocking
			ch <- 10 * i
		}
		close(ch)
	}()
	return ch
}
