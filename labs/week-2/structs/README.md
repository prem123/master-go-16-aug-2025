## Structs

Go supports what we call *compound* types, that is, types that are *composed* of other types. Structs allow you to group related data together, similar to objects in other programming languages.

These are called struct (for structure) types. We declare a struct like this:

```go
type Point struct {
  X int
  Y int
}
```

Point is a position in two dimensional space, it has two fields, `X` and `Y`.


## Define a Simple Struct

```go

// Define a struct to represent a Person
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    // Initialize a new instance of Person
    var person1 Person
    person1.FirstName = "John"
    person1.LastName = "Doe"
    person1.Age = 30

    // Print the struct
    fmt.Println(person1)
}

```
Run the program:
```bash
go run main.go
```

## Initialize Structs in Different Ways
Modify main.go to initialize structs in different ways:

```go

// Define a struct to represent a Person
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    // Option 1: Initialize using struct field assignments
    var person1 Person
    person1.FirstName = "John"
    person1.LastName = "Doe"
    person1.Age = 30

    // Option 2: Initialize using a struct literal
    person2 := Person{
        FirstName: "Alice",
        LastName:  "Smith",
        Age:       25,
    }

    // Option 3: Initialize using a struct literal without field names
    person3 := Person{"Bob", "Brown", 40}

    // Print all the structs
    fmt.Println(person1)
    fmt.Println(person2)
    fmt.Println(person3)
}
```

## Access and Modify Struct Fields
Update `main.go` to access and modify struct fields:

```go

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    person := Person{
        FirstName: "Emma",
        LastName:  "Johnson",
        Age:       28,
    }

    // Access struct fields
    fmt.Println("First Name:", person.FirstName)
    fmt.Println("Last Name:", person.LastName)
    fmt.Println("Age:", person.Age)

    // Modify struct fields
    person.Age = 29
    fmt.Println("Updated Age:", person.Age)
}
```

## Struct Embedding (Inheritance)
In Go, you can embed one struct inside another. This allows you to achieve a form of inheritance.

```go

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// Employee embeds Person struct
type Employee struct {
    Person
    EmployeeID string
}

func main() {
    // Initialize Employee struct with embedded Person
    employee := Employee{
        Person:     Person{"Sophia", "Miller", 30},
        EmployeeID: "E12345",
    }

    // Access fields from both Employee and embedded Person
    fmt.Println("Employee ID:", employee.EmployeeID)
    fmt.Println("Employee Name:", employee.FirstName, employee.LastName)
    fmt.Println("Employee Age:", employee.Age)
}

```

## Comparing Structs and Copying
Structs in Go are value types, meaning they are copied when assigned to a new variable or passed to a function. 

```go

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    // Create two Person instances
    person1 := Person{"Mia", "Taylor", 22}
    person2 := person1 // This creates a copy of person1

    // Modify the copy
    person2.Age = 23

    // Compare the two structs
    fmt.Println("Person 1:", person1)
    fmt.Println("Person 2 (modified):", person2)

    // Compare structs
    fmt.Println("Are person1 and person2 equal?", person1 == person2)
}

```

## String method

Any type that implements a `String() string` method will be used by the `fmt` package
when it prints the value.

