package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
	GPA  float64
}

type StudentManager struct {
	Students []Student
}

func (sm *StudentManager) AddStudent(s Student) {
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

func main() {
	sm := StudentManager{} // Create a new StudentManager instance

	// Example usage:
	// Add students
	sm.AddStudent(Student{Name: "Alice", Age: 20, GPA: 3.5})
	sm.AddStudent(Student{Name: "Bob", Age: 21, GPA: 3.2})
	sm.AddStudent(Student{Name: "Charlie", Age: 22, GPA: 3.9})

	// Display students
	fmt.Println("Students:")
	sm.DisplayStudents()

	// Update student
	sm.UpdateStudent("Bob", Student{Name: "Bob", Age: 22, GPA: 3.3})
	fmt.Println("\nAfter updating Bob's details:")
	sm.DisplayStudents()

	// Delete student
	sm.DeleteStudent("Charlie")
	fmt.Println("\nAfter deleting Charlie:")
	sm.DisplayStudents()
}
