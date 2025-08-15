## Formatting

Go provides excellent support for formatting using the `fmt.Printf` function. 

[fmt Documentation](https://pkg.go.dev/fmt)

📋 Here are some examples of common string formatting tasks:

```go
p := "New Delhi"
fmt.Printf("type: %T\n", p)

charA, charZ := 97, 122
fmt.Printf("type: %c and %c\n", charA, charZ)

fmt.Printf("str1: %s\n", "\"string\"")

fmt.Printf("str2: %q\n", "\"string\"")

fmt.Printf("width1: |%6d|%6d|\n", 19, 731)

fmt.Printf("width2: |%6.2f|%6.2f|\n", 6.7, 12.82)

fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
```

## Sprintf
📋  `Sprintf` formats and returns a string without printing it anywhere.

```go
char := 97
str := fmt.Sprintf("%c", char)
fmt.Println(str)

s := fmt.Sprintf("sprintf: a %s", "string")
fmt.Println(s)


// declaring variables of different datatypes
var message string = "Hello and welcome to"
var year int = 2022

// printing out the declared variables following the format string
var complete_msg = fmt.Sprintf("%s educative in %d", message, year)
fmt.Print(complete_msg)
```
