package main

import "fmt"

// Any struct that defines a receiver function named as FindAll is of Repository type
type Repository interface { // abstraction
	FindAll()
}

// This is real DB repository
type DBRepository struct{} // concrete implementation

// DBRepository is implementing Repository
func (db DBRepository) FindAll() {
	fmt.Println("I talk to DB and find information....")
}

// Fake repository
type StubRepository struct{} // concrete implementation

// StubRepository is implementing Repository
func (db StubRepository) FindAll() {
	fmt.Println("I am Fake/Stub and dont do anything....")
}

// type APIRepository struct{} // concrete implementation

// func (db APIRepository) FindAll() {
// 	fmt.Println("I am API repository....")
// }

func main() {

	db := DBRepository{}
	// I will handover the repository to getData and it will be responsible for getting data from repository
	getData(db)

	// I want the same getData function to also handle StubRepository
	stub := StubRepository{}
	getData(stub)

	printAnything(100)
	printAnything(200.20)
	printAnything("Hello World!!")
}

func getData(db Repository) {
	db.FindAll()
}

// func printAnything(data interface{}) { // you can pass any type
func printAnything(data any) { // any is an alias to interface{}
	fmt.Println(data)
}

type MyFileHandler struct{} // concrete implementation of Reader, Writer and ReadWriter from io package

func (m MyFileHandler) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (m MyFileHandler) Write(p []byte) (n int, err error) {
	return 0, nil
}
