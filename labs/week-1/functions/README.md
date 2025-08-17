## Functions
Go is all made up of functions, in fact, you've been writing functions all along.

You can declare your own functions with the func declaration.

A function's name must be a valid identifier, just like `const` and `var`.

📋 Creating a function `print`:

```go
func print(msg string) {
  fmt.Println(msg)
}

func main() {
  print("Hello World!!")
}
```

## Parameters
Parameters are defined in the function definition and the argument passed to the function should match the data type.

```go
// here the msg is the variable name and string is the datatype this function expects as a parameter
func print(msg string) {
  fmt.Println(msg)
}

hello(1234)
// compile this code and see that happens
```

## Variadic Functions in Go

A variadic function is a function that can accept a variable (changing) number of arguments.
This is useful when you don’t know in advance how many values will be passed.

- Declared using ... before the type of the last parameter.
```go
func myFunc(nums ...int) { }
```
- Inside the function, `nums` behaves like a slice of that type ([]int in this case).
- You can pass **zero, one, or many arguments**.
- You can also **expand an existing `slice`** into variadic arguments using `slice...`

### Example 1: Sum of Numbers

```go
// variadic function
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum())              // 0 (no numbers passed)
    fmt.Println(sum(1, 2, 3))       // 6
    fmt.Println(sum(10, 20, 30, 40)) // 100

    numbers := []int{5, 10, 15}
    fmt.Println(sum(numbers...)) // expand slice → 30
}
```

### Example 2: Print Any Number of Strings

```go
// variadic with strings
func greet(names ...string) {
    for _, n := range names {
        fmt.Printf("Hello, %s!\n", n)
    }
}

func main() {
    greet("Alice")
    greet("Bob", "Charlie", "Diana")
}
```

### Example 3: Mixing Fixed + Variadic Parameters

```go
// first argument is required, rest are optional
func introduce(role string, names ...string) {
    fmt.Printf("Role: %s\n", role)
    for _, n := range names {
        fmt.Println(" -", n)
    }
}

func main() {
    introduce("Instructor", "Ashish", "Neha")
    introduce("Student", "Rahul")
}

```

### 📝 Lab: Variadic Functions Practice

Problem Statement

Write a Go program with a variadic function that finds the maximum number among any number of integers.

### Requirements

1. Define a function `max(nums ...int) int` that:
- Accepts zero, one, or many integers.
- Returns the **largest integer**.
- If no numbers are passed, return 0 (or a message).

2. In `main()`:
- Call `max()` with different sets of numbers:
- No numbers
- Three numbers
- A `slice` expanded with ...

3. Print the result for each case.

### 📝 Bonus Challenge: 

Create another variadic function `concat(words ...string) string` that joins any number of words into a single sentence.

## Return value
The function definition should define the return type if the method returns a value.

```go
// Here last int is the return type
func sum(a int, b int) int {
  return a + b
}
```

One of Go's unusual features is that functions and methods can return multiple values. In case of multiple return type, you need to put them in parenthesis (n int, err error)

```go
// In Go, Write function can return a count and an error: "Yes, you wrote some bytes but not all of them because the disk was full".

func Write(b []byte) (n int, err error)
```

## Multiple return values
Multiple assignment is important to understand because you can return multiple values from a function.

```go
// This is a function declaration for sum, which takes two arguments, two int, and returns an int.
func sum(a int, b int) int
```

```go
// This is a function declaration for h, which takes two arguments, two ints, and returns three values, two int s and a string.
func h(i, j int) (int, int, string)
```

📋 Your program must return the number of values specified in the function signature.
```go
package main

import "fmt"

// swap switches the values of a and b
func swap(a int, b int) (int, int) {
  return b, a
}

func main() {
  a, b := swap(2, 5)
  fmt.Println(a, b)
}
```

When you call a function that returns multiple values, you must assign all of them or none
of them.

If you want to use only the first value you can ignore the second by assigning it
to the underscore variable, _.

```go
func main() {
  a, _ := swap(2, 5)
  fmt.Println(a)
}
```

## Named result parameters
The return or result "parameters" of a Go function can be given names. 

The names are not mandatory but they can make code shorter and clearer: they're documentation. If we name the results of `nextInt` it becomes obvious which returned `int` is which.

```go
func nextInt(b []byte, pos int) (value, nextPos int)
```
## Lab
Refactor the ["Interactive Calculator Lab"](../lab/readme.md) solution into functions. Ensure that the functions follow SRP

## Defer
Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns. It's an unusual but effective way to deal with situations such as resources that must be released regardless of which path a function takes to return. The canonical examples are unlocking a mutex or closing a file. 

## 📋 Example-1:

```go
func main() {
  defer releaseResources()
  fmt.Println("testing defer statement")
}

func releaseResources() {
  fmt.Println("closing and releasinng connections....")
}
```
## 📋 Example-2:

```go
func main() {
    fmt.Println("Start")

    // Defer this function call
    defer fmt.Println("Deferred: First")

    // Another defer
    defer fmt.Println("Deferred: Second")

    fmt.Println("End")
}
```

## 📋 Using defer for Resource Cleanup

```go
func main() {
    // Open a file
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }

    // Defer the closing of the file
    defer file.Close()

    // Perform some file operations
    fmt.Println("File opened successfully.")
}
```

Create a dummy file named test.txt in your project directory:
```bash
echo "Hello, Go defer!" > test.txt

# run the program
go run main.go
```

## 📋 Defer with Error Handling

In real-world applications, you may want to handle errors that occur during resource cleanup. Let’s update the file-handling example to include error checking.

```go

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }

    // Defer the closing of the file and check for errors
    defer func() {
        if err := file.Close(); err != nil {
            fmt.Println("Error closing file:", err)
        }
    }()

    // Simulate some file operations
    fmt.Println("File opened successfully.")
}
```
The deferred function now checks for errors when closing the file, ensuring proper handling in case the file closure fails.