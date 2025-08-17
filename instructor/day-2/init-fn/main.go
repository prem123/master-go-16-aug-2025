package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("main function invoked...")
}

func init() {
	log.Println("init invoked...")

	dburl := os.Getenv("MY_DB_CONN")
	if dburl == "" {
		log.Println("DB url not found, sanity check failed. exiting...")
		os.Exit(1)
	}
}
