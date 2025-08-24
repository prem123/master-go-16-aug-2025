package main

import (
	"fmt"
	"time"
)

type Point struct {
	X int
	Y int
}

type Customer struct {
	Id      string
	Name    string
	Address string
	City    string
}

type ShoppingCart struct {
	Item       // embedding
	Qty        int
	TotalPrice float64
}

type Item struct {
	Name       string
	MfdDate    time.Time
	ExpiryDate time.Time
}

func main() {
	p := Point{X: 10, Y: 20}

	// p1 := Point{10, 20}

	p2 := Point{Y: 20}

	fmt.Println("P: ", p)
	fmt.Println("P1: ", Point{10, 20})
	fmt.Println("P2: ", p2)

	fmt.Println("Empty Point: ", Point{})
}
