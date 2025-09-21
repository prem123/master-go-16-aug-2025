package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go f1(wg) // schedule this to run in future

	wg.Add(1)
	go f1(wg) // schedule this to run in future

	f2()

	wg.Wait()
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	fmt.Println("invoked f1")
}

func f2() {
	fmt.Println("invoked f2")
}
