package main

import "fmt"

func main() {

	// Go is "pass by value" in most of the cases and not in all the cases
	// whereever it is not passed by value, we need to explicit and that is
	// called "pass by reference"

	wallet := 2000

	fmt.Println("address in main: ", &wallet)

	fmt.Println("Before wallet amount: ", wallet)
	makePurchase(&wallet, 500) // here we are explictily specifying that we want to pass the reference by using '&' symbol
	fmt.Println("After wallet amount: ", wallet)
}

func makePurchase(wallet *int, purchaseAmount int) {
	fmt.Println("Inside make purchase: ", wallet)
	fmt.Println("value inside make purchase: ", *wallet) // dereferencing the address
	*wallet = *wallet - purchaseAmount
	// fmt.Println("while make purchase the final wallet amount: ", wallet)
}
