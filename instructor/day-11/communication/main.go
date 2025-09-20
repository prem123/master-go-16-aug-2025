package main

import (
	"fmt"
	"sync"
)

var result int

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go update(wg)

	wg.Add(1)
	go add(10, 20, wg)

	wg.Wait()
	fmt.Println("Result: ", result)
}

func update(wg *sync.WaitGroup) {
	defer wg.Done()
	result = 500
}

func add(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	result = a + b
}
