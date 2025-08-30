package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Lab 1: Sum of Letters using Maps ===")
	testSumOfLetters()
	
	fmt.Println("\n=== Lab 2: Word Size Analysis ===")
	analyzeWordSizes()
}

func testSumOfLetters() {
	testCases := []string{"", "a", "z", "cab", "excellent"}
	
	for _, testCase := range testCases {
		result := sumOfLetters(testCase)
		fmt.Printf("sumOfLetters(\"%s\") => %d\n", testCase, result)
	}
}

func sumOfLetters(letters string) int {
	letterValue := make(map[rune]int)
	
	for i, char := range "abcdefghijklmnopqrstuvwxyz" {
		letterValue[char] = i + 1
	}
	
	sum := 0
	for _, char := range letters {
		sum += letterValue[char]
	}
	return sum
}

func analyzeWordSizes() {
	str := `Commodo Lorem enim dolore culpa minim ipsum ex excepteur in Commodo duis nulla ex laborum irure sunt incididunt Incididunt amet Lorem amet dolor sit consectetur culpa esse quis laborum pariatur laborum fugiat mollit Mollit voluptate aliquip Lorem incididunt mollit pariatur eu enim proident culpa esse laborum voluptate Nostrud aliqua magna ipsum qui duis euNisi pariatur sit do magna Lorem nostrud voluptate occaecat occaecat quis dolore irure Velit aliqua reprehenderit eu duis aliqua excepteur duis non enim Nostrud qui voluptate enim eiusmod dolor proident laboris nostrud commodo laborum aliquip sunt Exercitation do anim do ullamco Ipsum eiusmod aute qui ea consectetur Veniam laborum occaecat mollit pariatur commodo id ullamco dolore ipsum sit dolore elit fugiatUt ad aliqua dolor quis velit reprehenderit Dolore aliquip exercitation do fugiat Pariatur irure aliqua magna quis`
	
	wordSizeCount := make(map[int]int)
	words := strings.Split(str, " ")
	fmt.Println("Words:", words)
	// fmt.Printf("Total words: %d\n", len(words))
	// fmt.Printf("First 10 words: %v\n", words[:10])
	
	for _, word := range words {
		size := len(word)
		wordSizeCount[size]++
	}
	
	var maxSize, maxCount int
	for size, count := range wordSizeCount {
		if count > maxCount {
			maxSize = size
			maxCount = count
		}
	}
	
	fmt.Printf("The highest occurring word size is: %d\n", maxSize)
	fmt.Printf("The number of occurrences for word size %d is: %d\n", maxSize, maxCount)
}
