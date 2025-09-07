package repository

import (
	"errors"
	"student-backend/models"
)

var sm = models.StudentManager{} // Create a new StudentManager instance

func init() {
	sm.Enroll(models.Student{Id: 1, Name: "Alice", Age: 20, GPA: 4.8, Password: "pass1"})
	sm.Enroll(models.Student{Id: 2, Name: "Bob", Age: 21, GPA: 3.2})
	sm.Enroll(models.Student{Id: 3, Name: "Charlie", Age: 22, GPA: 3.9})
}

// naming convention in repository
// Save			-> C
// FindAll	-> R
// FindById -> R
// Update		-> U
// Delete		-> D

func FindAll() []models.Student {
	return sm.Students
}

func FindById(id int) (*models.Student, error) {
	for _, student := range sm.Students {
		if student.Id == id {
			return &student, nil
		}
	}
	return nil, errors.New("id not found")
}

func Save(newStudent models.Student) int {
	newStudent.Id = len(sm.Students) + 1
	sm.Enroll(newStudent)
	return newStudent.Id
}
