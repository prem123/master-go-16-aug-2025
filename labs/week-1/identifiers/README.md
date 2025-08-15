## Identifier in Go

An identifier is a name your give to something in your code.

- In Go an identifier is any word that starts with a letter.
- A letter is anything that Unicode defines to be a letter, this includes Kanji, Cyrillic, Latin, etc.
- An identifier must start with a Unicode letter, or the underscore character, _.
- Numbers are not permitted at the start of an identifier, but can appear later in the identifier.

Open `main.go` and update the `main()` function:
```go
package main
import "fmt"

func main() {
   // creating an identifier using Unicode character
   var π = 3.14159
   fmt.Println(π)
}
```

## Multiple assignment

Go allows you to perform multiple assignments and declarations in one statement.

📋 For example, if we wanted to declare, x, y, and z, with the values 1, 2, and 3 respectively.
We could write:

```go
var x = 1
var y = 2
var z = 3
```

We can write the same thing like this:

```go
var x, y, z = 1, 2, 3
```

## Declarations
There are six kinds of declarations in Go, we've seen most of them already:

- `const`: declares a new constant.
- `var`: declares a new variable.
- `type`: declares a new type.
- `func`: declares a new function, or method.
- `package`: declares the package this .go source file belongs to.
- `import`: declares that this package imports declarations from another.

## Variables
A variable holds a value that can be changed over time. You declare a new variable with the var declaration.

Just like constants, variable identifiers can be any any valid Unicode word.

## Unused variable declarations | Blank identifier
Unused variables are often the source of bugs. If you declare a variable in the scope of
your function but do not use it, the Go compiler will complain.

If the variable is unused, you should delete it, or assign it temporarily to the magic variable
called `_`

`_` in Go is also referred as [blank identifier](https://go.dev/doc/effective_go#blank). The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly.

```go
package main

import "fmt"

func main() {
  // not using below two variable will result in compilation error
  var pi = 3.14159
  var totalAplha = 26

  //uncommenting the below two lines will fix the compilation error

  //fmt.Println(pi)
  //fmt.Println(totalAlpha)

  //The value of pi and totalAlpha can also be assigned to _
  //_ = pi
  //_ = totalAplha

}

```

## Increment and decrement

Go supports a limited form of variable post-increment and post-decrement, ie. x++, x--.

```go
var i = 1
i++
println(i)
```

`i++` and `i--` are *statements*, not an *expressions*, they do not produce a value.

```go
var i = 1
var j = i++
println(i, j)
```
This program will result in a syntax error.