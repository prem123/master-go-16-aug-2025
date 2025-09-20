package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go getNos(ch) //scheduled for future

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func getNos(ch chan int) {
	time.Sleep(500 * time.Millisecond) //blocking
	ch <- 10
	time.Sleep(500 * time.Millisecond) //blocking
	ch <- 20
	time.Sleep(500 * time.Millisecond) //blocking
	ch <- 30
	time.Sleep(500 * time.Millisecond) //blocking
	ch <- 40
	time.Sleep(500 * time.Millisecond) //blocking
	ch <- 50
}
