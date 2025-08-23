# 🧑‍💻 Week 1 Projects – Go Fundamentals

Welcome to Week 1 of your Go learning journey! Below are three hands-on projects designed to help you apply what you’ve learned. These projects reinforce variables, loops, conditionals, functions, type safety, and I/O.

---

## ⏱ 1. GoTimer – Countdown Timer App

### 📝 Problem Statement
Create a terminal-based countdown timer that accepts a number of seconds from the user and counts down to zero.

### Objective
Use loops, basic time handling, and functions to build a CLI countdown experience.

### What You'll Build
- User inputs time in seconds
- The program counts down and prints each second
- Prints a final message when the time is up
- User can start another timer or exit

### Key Features
- Use `fmt.Scanln` to read input
- Loop from N to 1 using `for`
- Sleep 1 second per iteration using `time.Sleep`
- Use `defer` to show a goodbye message
- Use a custom function to notify when the timer ends
- (Bonus) Use function passing for the alert message

<details>
  <summary>Not sure how?</summary>

```go
func main() {
	defer fmt.Println("Goodbye! Thanks for using the timer.") // defer demo

	for {
		var seconds int
		fmt.Print("Enter number of seconds for countdown (0 to exit): ")
		fmt.Scanln(&seconds)

		if seconds <= 0 {
			return // exit program
		}

		// run the countdown
		runCountdown(seconds, alertMessage)
	}
}

// runCountdown handles the countdown logic
func runCountdown(seconds int, alert func()) {
	for i := seconds; i > 0; i-- {
		fmt.Printf("⏳ %d\n", i)
		time.Sleep(1 * time.Second)
	}
	// call the passed function
	alert()
}

// alertMessage is the function to notify when timer ends
func alertMessage() {
	fmt.Println("⏰ Time’s up!")
}

```  
</details>
<br>

---

## 2. Quiz Me! – A Simple CLI Quiz Game

📝 You’ll build a small command-line quiz program that asks the user a few questions, checks their answers, and shows a score at the end.

### Guidelines

- Keep the questions, options and the correct answer in slice.
- Represent them as a `slice` of slices (`[][]string`).
- Use a `for` loop to go through each question. 
- Accept both uppercase and lowercase answers (`A` or `a`).
- Print a message `"Thanks for playing!"` when the program finishes, no matter how it ends.

### Functions

- Create a function `askQuestion(num int, q []string) (bool, error)`:
  - Prints the question.
  - Reads the user’s input.
  - Returns `true` if correct, `false` otherwise.
  - Demonstrates multiple return values (`bool` + `error`).

- Create another function `calculateScore(score int, correct bool) int`:
  - Updates the score if the answer is correct.
  - Keeps your code clean and modular.


```bash
$ go run main.go
Welcome to Quiz Me!

Q1: What is the keyword to declare a constant in Go?
a) const
b) var
c) constant
Your answer: a
Correct!

Q2: Which of these is zero value of int in Go?
a) 0
b) nil
c) undefined
Your answer: b
Wrong! Correct answer is a

You scored 1/2
Thanks for playing!

```
---


<details>
  <summary>Not sure how?</summary>

```go
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Thanks for playing!") // defer demo

	questions := [][]string{
		{"What is the keyword to declare a constant in Go?",
			"const", "var", "constant", "a"},
		{"Which of these is zero value of int in Go?",
			"0", "nil", "undefined", "a"},
	}
	score := 0

	for i, q := range questions {
		correct, _ := askQuestion(i+1, q)
		score = calculateScore(score, correct)
	}

	fmt.Printf("Final Score: %d/%d\n", score, len(questions))
}

func askQuestion(num int, q []string) (bool, error) {
	fmt.Printf("\nQ%d: %s\n", num, q[0])
	fmt.Printf("a) %s\nb) %s\nc) %s\n", q[1], q[2], q[3])
	fmt.Print("Your answer: ")

	var input string
	fmt.Scanln(&input)

	if input == q[4] {
		fmt.Println("Correct!")
		return true, nil
	}
	fmt.Println("Wrong!")
	return false, nil
}

func calculateScore(score int, correct bool) int {
	if correct {
		return score + 1
	}
	return score
}

```
</details>
<br>