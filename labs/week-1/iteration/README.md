## Looping / Iteration
Go has a single for loop construct that combines
- `while` condition { … }
- `do { … } while` condition
- `do { … } until` condition

into one syntax.

- `for (init statement); condition; (post statement) { … }`

The parts of a for statement are:
- `init statement`: used to initalise the loop variable; i = 0.
- `condition`: user to test if the loop is done; i < 10, true means keep looping.
- `post statement`: user to increment the loop variable; i++, i = i - 1.

Let's write a small program using loop:

Keep the below code in a new file named as `main.go` inside `loop` folder. You have to create the folder

```go
package main

import "fmt"

func main() {
  var i = 0
  for i = 1; i < 11; i++ {
    println(i)
  }
}
```

Note: No need to put ( braces around the for condition ). In fact, if you do it's a syntax error. Try it.

Now let's try using the for loop as a while in other languages:

```go
  var i = 10
  for i > 0 {
    println(i)
    i--
  }
```
Edit the program to make it print only the numbers from 6 down to 4.

Note: This for loop only has a condition, there is no init statement or post statement, so we
can omit the semicolons ;

**Printing alphabets from `a` to `z`:**

```go
// Go uses the `ASCII` codes within the `for` loop to print alphabets from `a` to `z`.
for ch := 'a'; ch <= 'z'; ch++ {
  fmt.Printf("%c ", ch)
}
```

## Lab:

Using loop in Go, print the following structure

```
*
* *
* * *
* * * *
* * * * *
```