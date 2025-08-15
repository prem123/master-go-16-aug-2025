## Control statements (conditional)

Go has two conditional statements, *`if`* and *`switch`*.

```go
if v > 0 {
    println("v is greater than zero")
} else {
    println("v is less than or equal to zero")
}
```

Let's write a program to print even numbers between 1 to 10

<details>
  <summary>Not sure how?</summary>

Create a file `main.go` inside folder named as `conditional` and add the code:

```go
package main

import "fmt"

func main() {
  var i = 0
  for i = 1; i < 11; i++ {
    if i % 2 == 0 {
      fmt.Println(i)
    }
  }
}
```
</details>
<br>

## continue
`continue` is used to skip the body of the loop and start from top.

```go
var i = 0
  for i = 1; i < 11; i++ {
    // If statements in Go are often used as guard clauses.
    if i % 2 == 1 {
      continue
    }
  }
  fmt.Println(i)
}
```

## break
`break` statement can be used to break out of the current loop

```go
var i = 1
for {
  if i > 10 {
    break
  }
  if i%2 == 0 {
    fmt.Println(i)
  }
  i++
}
```

## switch
The switch statement lets you check multiple cases. You can see this as an alternative to a list of if-statements. A switch statement provides a clean and readable way to evaluate cases.

```go
var i = 2
switch i {
  case 1:
    println("value is one")
  case 2:
    println("value is two")
  default:
    println("default case")
}
```
you can also check conditions:

```go
x := 5
switch {
  case x < 10:
    f1() //calling function f1
  case x > 10:
    f2() //calling function f2
  default:
    println("default called")
}

func f1() {
  fmt.Println("f1 called")
}

func f2() {
  fmt.Println("f2 called")
}
```

There is also a third form, it includes an initialization statement:
```go
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

<br>

## Lab 1:

In the last unit test example you have seen the `addContact` function having a chaining of `if-else` for phone number. Convert the `if-else` chaining to `switch`

[Unit test program](../unit-testing#write-unit-test)


## Lab 2:

1. Write a function called `hello` that receives a name and prints hello, <name>
for e.g `Hello("Chris")` prints `Hello, Chris`

2. If the name is passed as empty, it should print "Hello, World", e.g. `Hello("")`

3. Modify the Hello function to receive one more parameter called language `Hello(name string, lang string)` and should support `Spanish` as a language. Program should print:
    - `Hola, Chris` when `Hello("Chris", "Spanish")`
    - `Hello, Chris` when `Hello("Chris", "")`
    - `Hello, World` when `Hello("", "")`

4. Add a new requirement to support one more language as `French`. Program should print:
    - `Bonjour, Chris` when `Hello("Chris", "French")`.
    - rest of the requirement will remain same.
    - modify the program to use `switch` statement instead of using `if`.

5. Write unit tests