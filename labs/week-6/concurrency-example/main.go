package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var max int = 10e7
	calculateSumOfNumbers(max)
	// calculateConcurrent(max)
}

func calculateConcurrent(max int) {
	s := GenerateNumbers(max)
	t := time.Now()
	ch := make(chan int64)

	cpuCount := runtime.NumCPU()
	// cpuCount := 2
	// runtime.GOMAXPROCS(cpuCount)
	sizeOfParts := max / cpuCount

	for i := 0; i < cpuCount; i++ {
		start := i * sizeOfParts
		end := start + sizeOfParts

		part := s[start:end]
		go sum(part, ch)
	}
	var total int64
	for i := 0; i < cpuCount; i++ {
		total = total + <-ch
	}
	fmt.Printf("Concurrent ALL Cpu Add, Sum: %d,  Time Taken: %s\n", total, time.Since(t))

}

func calculateSumOfNumbers(max int) {
	s := GenerateNumbers(max)
	t := time.Now()
	ch := make(chan int64)

	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	x, y := <-ch, <-ch
	total := x + y
	fmt.Printf("Without channel Add, Sum: %d,  Time Taken: %s\n", total, time.Since(t))
}

func sum(s []int, ch chan int64) {
	var sum int64
	for _, v := range s {
		sum += int64(v)
	}
	ch <- sum
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}
