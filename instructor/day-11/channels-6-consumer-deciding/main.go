package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	count := 10

	go getNos(ch, count) //scheduled for future

	// consumer
	for i := 1; i <= count; i++ {
		fmt.Println(<-ch)
	}
}

// producer
func getNos(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		time.Sleep(500 * time.Millisecond) //blocking
		ch <- 10 * i
	}
}
