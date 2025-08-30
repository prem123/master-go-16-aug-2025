package main

import "fmt"

func main() {
	fmt.Println("=== Lab 1: Subslices ===")
	lab1()
	
	fmt.Println("\n=== Lab 2: Sum of Letters ===")
	lab2()
}

// Lab 1: Print primes less than 10 and two-digit primes using subslices
func lab1() {
	// These are the primes less than 200
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109,
		113, 127, 131, 137, 139, 149, 151, 157,
		163, 167, 173, 179, 181, 191, 193, 197, 199}
	
	fmt.Println("All primes:", primes)

	// Print all primes less than 10
	fmt.Print("Primes less than 10: ")
	for i := 0; i < len(primes); i++ {
		if primes[i] >= 10 {
			primesLessThan10 := primes[0:i]
			fmt.Println(primesLessThan10)
			break
		}
	}

	// Bonus: Print only the two-digit primes
	var startIdx, endIdx int
	
	// Find start index (first prime >= 10)
	for i := 0; i < len(primes); i++ {
		if primes[i] >= 10 {
			startIdx = i
			break
		}
	}
	
	// Find end index (first prime >= 100)
	for i := startIdx; i < len(primes); i++ {
		if primes[i] >= 100 {
			endIdx = i
			break
		}
	}
	
	twoDigitPrimes := primes[startIdx:endIdx]
	fmt.Println("Two-digit primes:", twoDigitPrimes)
}



func sumOfLetters(letters string) int {
	sum := 0
	for _, char := range letters {
		if char >= 'a' && char <= 'z' {
			value := int(char - 'a' + 1)
			sum += value
		}
	}
	return sum
}
