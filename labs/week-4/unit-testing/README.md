## Unit Testing

Testing is built in to the language and the syntax is same as you write rest of the code.

Writing a test is just like writing a function, with a few rules:
- Test needs to be in a file with a name like `xxx_test.go`
- The test function must start with the word `Test`
- The test function takes one argument only `t *testing.T`
- In order to use the `*testing.T` type, you need to import `"testing"` package

For now, it's enough to know that `t` of type `*testing.T` is your "hook" into the testing framework so you can do things like `t.Fail()` or `t.Error()` when you want to fail.

A simple test in Go:

```go
package main

import (
  "testing"
)

func TestGreeting(t *testing.T) {
  got := greeting()
  want := "Hello"
  if got != want {
    t.Errorf("greeting = %s; wanted Hello", got)
  }
}

func greeting() string {
  return "Hello"
}
```

To run test:
```shell
$ go test
```
You might get the modules error depending on your version of go:

```
go: go.mod file not found in current directory or any parent directory; see 'go help modules'
```

This can be easily fixed by creating a go module. More on modules later:

```shell
$ go mod init unit-testing-example

# it might ask you to run the tidy command
$ go mod tidy

# run tests again
$ go test
```

## Test coverage
`go test` can report coverage

```shell
go test -coverprofile=cover.out
```

This produces a coverage file, `cover.out`:

- `go tool cover -func=cover.out` will print the coverage report
- `go tool cover -html=cover.out` will open the report in a browser

<br>

## Lab: Write unit test
Write unit tests for below program:

`addContact` function add a contact, before adding it validates the input parameters.

*Rules:*

- first name cannot be blank.
- last name cannot be blank.
- phone number cannot blank.
- phone number length should be more than 10.
- phone number should contain only digits.
- phone number should start with 0.

When you are done, also check the coverage of your program.

```go
package main

import (
  "errors"
  "fmt"
  "strconv"
)

func main() {
  err := addContact("first", "last", "phone")
  if err != nil {
    fmt.Println("error while adding contact: " + err.Error())
  } else {
    fmt.Println("contact added successfully")
  }
}

func addContact(firstName string, lastName string, phone string) error {
  if firstName == "" {
    return errors.New("first name cannot be blank")
  }

  if lastName == "" {
    return errors.New("last name cannot be blank")
  }

  if phone == "" {
    return errors.New("phone cannot be blank")
  } else {
    if len(phone) < 10 {
      return errors.New("phone number should be of atleast 10 digits")
    }
    firstChar := phone[0:1]
    if firstChar != "0" {
      return errors.New("phone number should start with zero")
    } else {
      _, err := strconv.Atoi(phone)
      if err != nil {
        return errors.New("phone number should only be numeric: " + err.Error())
      }
    }
  }
  return nil
}
```
