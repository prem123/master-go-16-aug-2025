# 🧑‍💻 Week 1 Projects – Go Fundamentals

Welcome to Week 1 of your Go learning journey! Below are three hands-on projects designed to help you apply what you’ve learned. These projects reinforce variables, loops, conditionals, functions, type safety, and I/O.

---

## 🧠 1. GoQuiz – Command-Line Quiz App

### 📝 Problem Statement
Build a command-line quiz application that asks the user a series of questions, takes their input, and tracks their score. Add a time limit to complete the quiz.

### 🎯 Objective
Use Go syntax, conditionals, loops, and functions to build an interactive quiz app that responds to user input and provides feedback.

### ✅ What You'll Build
- A 5-question quiz in the terminal
- Score tracking and final result
- A global timer that ends the quiz after 30 seconds
- Answer validation using functions

### 🧱 Key Features
- Store questions/answers in a slice
- Read input using `bufio.Scanner`
- Use `if/else` to check answers
- Show final score using `defer`
- Use `select` and `time.After` to implement quiz timeout
- (Bonus) Use a function to validate answers

### 🧾 Deliverables
- `goquiz/main.go`
- Final output includes score and a farewell message

---

## ➕ 2. GoCalc – Command-Line Calculator

### 📝 Problem Statement
Build a simple calculator that takes two numbers and an arithmetic operator from the user and prints the result.

### 🎯 Objective
Apply core Go fundamentals to create a reusable and interactive CLI calculator.

### ✅ What You'll Build
- CLI calculator that supports `+`, `-`, `*`, `/`
- Input validation and looping
- Multiple operations in a single session

### 🧱 Key Features
- Use `fmt.Scanln` or `bufio.Scanner` to read user input
- Parse and validate numeric input
- Use `switch` or `if/else` to handle different operations
- Define separate functions for each operation
- (Bonus) Use first-class functions to map operators

### 🧾 Deliverables
- `gocalc/main.go`
- Working calculator with repeat option and error handling

---

## ⏱ 3. GoTimer – Countdown Timer App

### 📝 Problem Statement
Create a terminal-based countdown timer that accepts a number of seconds from the user and counts down to zero.

### 🎯 Objective
Use loops, basic time handling, and functions to build a CLI countdown experience.

### ✅ What You'll Build
- User inputs time in seconds
- The program counts down and prints each second
- Prints a final message when the time is up
- User can start another timer or exit

### 🧱 Key Features
- Use `bufio.Scanner` to read input
- Loop from N to 1 using `for`
- Sleep 1 second per iteration using `time.Sleep`
- Use `defer` to show a goodbye message
- Use a custom function to notify when the timer ends
- (Bonus) Use function passing for the alert message

### 🧾 Deliverables
- `gotimer/main.go`
- Fully working timer with clean input/output and graceful exit

---
