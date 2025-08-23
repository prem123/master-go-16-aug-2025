package main

import (
	"fmt"
	"os"
)

func main() {

	// Whenever we deal with resources, below are the steps
	// 1. allocate resource
	// 2. use/consume
	// 3. release

	// Database:
	// 1. Open connection
	// 2. use connection
	// 3. release connection to pool

	// File handling:
	// 1. open files
	// 2. file handling (read/write)
	// 3. close files

	// open file
	defer closeFile()                                  // this will be called last
	defer fmt.Println("going out of the function....") // this will be called second

	defer func() {
		fmt.Println("deferring a lot of steps") // this is first
	}()

	fmt.Println("Open file...")
	fmt.Println("File handling reading and writing from file...")
	// file handling

	mainFileFunc()
}

func mainFileFunc() {
	// Open a file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Defer the closing of the file
	defer func() {
		fmt.Println("closing file...")
		file.Close()
	}()

	// Perform some file operations
	fmt.Println("File opened successfully.")
}

func closeFile() {
	fmt.Println("I close file...")
}
