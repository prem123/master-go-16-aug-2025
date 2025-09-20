## Panic and Recover

- Starting with code:

```go
package main

import "fmt"

func riskyDivision(a, b int) {
	// defer + recover will handle the panic if b == 0
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Printf("Dividing %d by %d...\n", a, b)
	result := a / b // This will cause panic if b == 0
	fmt.Println("Result:", result)
}

func main() {
	fmt.Println("Program started")

	riskyDivision(10, 2) // Safe division
	riskyDivision(5, 0)  // Will cause panic but recover

	fmt.Println("Program continues after recovery")
}

```

## When to use `panic` and `recover`

1. ### Truly unrecoverable errors
- Something has gone so wrong that the program `cannot safely continue`.
- **Examples**:
  - Corrupted memory state
  - Impossible invariants being broken (e.g., index calculation logic bug)
  - System resource failure (e.g., failure to start a goroutine scheduler, library invariants violated)

2. ### Programming errors / Developer mistakes
- You hit a situation that should never happen if the code is correct.
- **Example**:
  ```go
  func mustPositive(n int) {
      if n <= 0 {
          panic("n must be positive")
      }
  }
  ```

3. ### Wrapping third-party panics
- When you’re building libraries or frameworks, you may want to `recover` to prevent panics from bubbling up and crashing the whole application.
- **Example**: Web servers (like Gin, Echo) use `recover` in middleware to catch user handler panics, log them, and return an `HTTP 500` instead of crashing the server.

4. ### Testing and quick prototypes
- Sometimes you want to crash fast to catch bugs early during development.
- **Example**: in `init` functions if critical configuration is missing.

## When not to use panic/recover

- Don’t use it for normal error handling (file not found, invalid input, network timeouts).
- In Go, the idiomatic way is:
  ```go
  data, err := os.ReadFile("config.json")
  if err != nil {
      return err // not panic
  }
  ```
- Don’t hide panics silently unless you log or propagate them. Swallowing them makes debugging very hard.

### Rule of Thumb

- Use `error` for expected issues.
- Use `panic/recover` only for unexpected, unrecoverable issues.