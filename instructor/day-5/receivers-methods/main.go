package main

import (
	"fmt"
)

//functions
// methods -> associated with Objects in other languages
// encapsulation
// in Go we call methods as receiver functions

type Inventory struct {
	ProductId   int
	ProductName string
	ProductQty  int
	Price       float64
}

func main() {
	item1 := Inventory{
		ProductId:   1,
		ProductName: "Mouse",
		ProductQty:  20,
		Price:       1000.00,
	}
	item2 := Inventory{
		ProductId:   2,
		ProductName: "Keyboard",
		ProductQty:  30,
		Price:       2000.00,
	}
	item3 := Inventory{
		ProductId:   3,
		ProductName: "Monitor",
		ProductQty:  10,
		Price:       5000.00,
	}

	// CRUD operations
	// C - Create/add more items to the inventory
	// R - Search the inventory
	// U - I can update items in the inventory
	// D - Delete items
	listOfItems := []Inventory{item1, item2, item3}

	// I want to iterate over all items and print the total price of the items?

	grandTotal := 0.0
	for _, item := range listOfItems {
		// total := calculateTotal(item.ProductQty, item.Price)

		total := item.Total() // this will be a receiver function

		grandTotal = grandTotal + total

		fmt.Println()
		fmt.Println(item)
		fmt.Println("Updating to new quantity by reducing 2")
		item.UpdateQty(item.ProductQty - 2)
		fmt.Println("After updating quantity...")
		fmt.Println(item)

	}
	fmt.Println("Grand Total: ", grandTotal)
}

// function
func calculateTotal(qty int, price float64) float64 {
	return float64(qty) * price
}

// behaviour is defined by the public functions of a class
// in case of Go, behaviour is defined by the "Exported receiver functions" of a struct

// receiver function
func (inv Inventory) Total() float64 {
	return float64(inv.ProductQty) * inv.Price
}

func (inv Inventory) Print() {
	fmt.Printf("ID: %d, Product Name: %15s, Total: %.2f\n", inv.ProductId, inv.ProductName, inv.Total())
}

// Changing the default implementation of String function
func (inv Inventory) String() string {
	return fmt.Sprintf("ID: %d, Qty: %d, Product Name: %15s, Total: %.2f", inv.ProductId, inv.ProductQty, inv.ProductName, inv.Total())
	// return fmt.Sprintf("{id: %d, productName: %15s, Total: %.2f}", inv.ProductId, inv.ProductName, inv.Total())
}

// This function will work only with the pointer because we are mutating the object
func (inv *Inventory) UpdateQty(newQty int) {
	inv.ProductQty = newQty
}
