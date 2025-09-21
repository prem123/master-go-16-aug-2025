package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // schedule this to run in future
	f2()
	fmt.Scanln()
}

func f1() {
	time.Sleep(5 * time.Second)
	fmt.Println("invoked f1")
}

func f2() {
	fmt.Println("invoked f2")
}
