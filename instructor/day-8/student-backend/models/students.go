package models

import "fmt"

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
