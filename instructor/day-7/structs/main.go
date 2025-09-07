package main

import (
	"fmt"
)

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {

	wallet := 2000
	fmt.Println(&wallet)

	// fmt.Println(wallet)
	// makePurchase(&wallet)
	// fmt.Println(wallet)

	emp := &Employee{1, "Some emp1", 5000}

	fmt.Println("Salary Before: ", emp.Salary)
	awardBonus(emp)
	fmt.Println("Name: ", emp.Name)
	fmt.Println("Salary After: ", emp.Salary)

	// fmt.Printf("%p\n", emp)
}

func awardBonus(emp *Employee) {
	// you dont need to use the dereferencing operator in case of structs. This make things easy
	// (*emp).Salary += 1000
	emp.Salary += 1000
}

func makePurchase(wallet *int) {
	*wallet = *wallet - 100
}
