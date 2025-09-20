package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go getNos(ch) //scheduled for future

	// consumer
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		break
	}
}

// producer
func getNos(ch chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond) //blocking
		ch <- 10 * i
	}
	close(ch)
}
