## Constants
Constants are values that do not change. There are six kinds of constants
- integer constants: 1, 0, -6, 9999999999999999999, …
- floating point constants: 3.14, 7.5e-2, …
- complex number constants (rare): 1 - 0.707i, …
- string constants: "Hello", …
- rune constants: 'a', 'す', '1', …
- boolean constants: true, false

To make a constant, we declare one with the `const` keyword

```
const msg = "Welcome to Go!!"
```

📋 Create the above constant in your Go program and print its value.
<details>
  <summary>Not sure how?</summary>

Open `main.go` and update the `main()` function:
```go
package main

import "fmt"

func main() {
   const msg = "Welcome to Go!"
   fmt.Println(msg)
}
```

</details>
<br>

Try changing the constant:
```go
func main() {
   const msg = "Welcome to Go!"
   fmt.Println(msg)
   //try changing msg. See that happens when you run
   msg = "overwhelming welcome!!"
}
```