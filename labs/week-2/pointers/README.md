## Pointers

Go supports pointers, allowing you to pass references to values and records within your program. Pointers allow you to reference memory locations, enabling efficient manipulation of data without copying values.

Example-1: 

```go

func main() {
    // Declare a variable
    a := 42

    // Declare a pointer that stores the address of 'a'
    var p *int = &a

    fmt.Println("Value of a:", a)
    fmt.Println("Address of a (pointer p):", p)
    fmt.Println("Value of a via pointer p:", *p) // Dereferencing the pointer to get the value

    // Modify the value of 'a' through the pointer
    *p = 100
    fmt.Println("New value of a (after modifying via pointer):", a)
}

```

Example-2: 

```go
package main

import "fmt"

func main() {
	wallet := 1000

	makePurchase(wallet, 200)
	fmt.Println("New wallet value:", wallet)
}

func makePurchase(wallet int, purchaseAmount int) {
	wallet = wallet - purchaseAmount
}
```

By default, arguments are passed by value. `makePurchase` will get a copy of wallet and that will be distinct from the one in the calling function.

This program can be fixed by using pointers:

```go
package main

import "fmt"

func main() {
	wallet := 1000

	makePurchase(&wallet, 200)
	fmt.Println("New wallet value:", wallet)
}

func makePurchase(wallet *int, purchaseAmount int) {
	*wallet = *wallet - purchaseAmount
}
```
After making the changes, `makePurchase` is receiving `wallet` as `pointer`. 

`&wallet` in the `main` function gives the memory address of wallet, i.e. a `pointer to wallet`.