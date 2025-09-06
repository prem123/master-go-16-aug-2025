package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`  // student_name
	Age      int     `json:"age"`   // age
	GPA      float64 `json:"score"` // score
	Password string  `json:"-"`     // hiding from json marshaller
}

type StudentManager struct {
	Students []Student // composition
}

func (sm *StudentManager) Enroll(s Student) {
	sm.Students = append(sm.Students, s)
	// Implement the logic to add a student to the list.
}

func (sm StudentManager) DisplayStudents() {
	for _, s := range sm.Students {
		fmt.Println(s)
	}
	// Implement the logic to display all students and their details.
}

/*
Modifying Elements of the Slice: In UpdateStudent function, we're modifying the
fields of the Student struct within the slice sm.Students. This doesn't affect the slice itself, but only the elements within it. Since you're modifying the elements, and not the slice itself, passing the receiver by value works fine.

When a slice is passed by value, only a copy of the slice header is made, not the underlying array. This slice header contains information about the length, capacity, and a pointer to the underlying array. Therefore, modifying elements of the slice (e.g., changing the values of elements) within the function will indeed reflect changes outside the function because both the original and the copy of the slice point to the same underlying array.
*/
func (sm StudentManager) UpdateStudent(name string, updatedStudent Student) {
	// Implement the logic to update a student's details by name.
	for i, s := range sm.Students {
		if s.Name == name {
			sm.Students[i].Age = updatedStudent.Age
			sm.Students[i].GPA = updatedStudent.GPA
			break
		}
	}
}

/*
Modifying the Slice Itself (Adding/Removing Elements): In our DeleteStudent function, we're modifying the slice itself by removing an element. In this case, passing the receiver by value creates a copy of the StudentManager struct, including a copy of the Students slice. If you modify this copy of the slice (e.g., by removing an element), it only affects the copy, not the original StudentManager instance.

If the receiver struct is passed by value: Modifying the slice itself (e.g., adding or removing elements) within the function will only affect the copy of the slice, not the original slice outside the function. This is because the copy of the slice header points to the same underlying array, but it's still a distinct slice with its own length and capacity metadata.

So, the difference in behavior comes from how the function interacts with the slice header, which contains the metadata about the slice, including the pointer to the underlying array. Passing the receiver struct by pointer allows the function to directly manipulate the original slice, including its metadata, while passing it by value creates a copy of the slice header, leading to separate metadata for the copy.
*/
func (sm *StudentManager) DeleteStudent(name string) {
	// Implement the logic to delete a student by name.
	for i, s := range sm.Students {
		if s.Name == name {
			sm.Students[i] = sm.Students[len(sm.Students)-1]
			sm.Students = sm.Students[:len(sm.Students)-1]
			break
		}
	}
}

var sm = StudentManager{} // Create a new StudentManager instance

// exercise for the day
// implement this: http://localhost:8080/students?id=1
func main() {
	// Example usage:
	// Add students
	sm.Enroll(Student{Name: "Alice", Age: 20, GPA: 3.5, Password: "pass1"})
	sm.Enroll(Student{Name: "Bob", Age: 21, GPA: 3.2})
	sm.Enroll(Student{Name: "Charlie", Age: 22, GPA: 3.9})

	// registering the end-point/route (/greet) with the request multiplexer
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/students", studentHandler)

	log.Println("Starting server on 8080 port....")
	http.ListenAndServe("localhost:8080", nil)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	// byte array -> struct

	// var num int
	// num := 0

	// var student Student
	student := Student{} // creating an empty new instance

	json.NewDecoder(r.Body).Decode(&student)

	// body, _ := r.GetBody()
	// var p []byte
	// body.Read(p)
	// json.Unmarshal(p, &student)
	/*
		{"name":"Alice","age":20,"score":3.5}
	*/
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	// charlie := Student{Name: "Charlie", Age: 22, GPA: 3.9}

	// struct -> byte array

	w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(sm.Students)

	bytes, _ := json.Marshal(sm.Students)

	// w.WriteHeader(200) // 201 is when we create new content
	w.Write(bytes)
}

//https://www.google.com/search?q=http+request+types&oq=http+request+types&gs_lcrp=EgZjaHJvbWUyBggAEEUYOTIGCAEQLhhA0gEINjU3M2owajGoAgCwAgA&sourceid=chrome&ie=UTF-8

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	msg := fmt.Sprintf("Hello %s, Welcome to Go Api", name)
	w.Write([]byte(msg))
}

func cui() {
	sm := StudentManager{} // Create a new StudentManager instance

	// Example usage:
	// Add students
	sm.Enroll(Student{Name: "Alice", Age: 20, GPA: 3.5})
	sm.Enroll(Student{Name: "Bob", Age: 21, GPA: 3.2})
	sm.Enroll(Student{Name: "Charlie", Age: 22, GPA: 3.9})
	fmt.Println("Students Manager:")
	// sm.DisplayStudents()

	input := 0
	for {
		fmt.Println("1. Enroll")
		fmt.Println("2. Display Students")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			newStudent := Student{}

			fmt.Print("Enter name: ")
			fmt.Scanln(&newStudent.Name)
			fmt.Print("Enter Age: ")
			fmt.Scanln(&newStudent.Age)
			fmt.Print("Enter GPA: ")
			fmt.Scanln(&newStudent.GPA)

			sm.Enroll(newStudent)
		case 2:
			sm.DisplayStudents()
		case 3:
			goto EXIT
		}
	}

EXIT:
	fmt.Println("Program exit...")

}
