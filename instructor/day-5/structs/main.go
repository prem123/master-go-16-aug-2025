package main

import "fmt"

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
		ProductQty:  2,
		Price:       1000.00,
	}
	item2 := Inventory{
		ProductId:   2,
		ProductName: "Keyboard",
		ProductQty:  3,
		Price:       2000.00,
	}
	item3 := Inventory{
		ProductId:   3,
		ProductName: "Monitor",
		ProductQty:  1,
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
		total := item.Price * float64(item.ProductQty)
		grandTotal = grandTotal + total
		fmt.Printf("ID: %d, Product Name: %15s, Total: %.2f\n", item.ProductId, item.ProductName, total)
	}
	fmt.Println("Grand Total: ", grandTotal)
}
