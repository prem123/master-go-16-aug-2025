package service

import (
	"os"
	"strconv"
	"student-backend/models"
	"student-backend/repository"
)

var TOP_GPA float64

func init() {
	gpa := os.Getenv("TOP_GPA")
	if gpa == "" {
		TOP_GPA = 4.5
	} else {
		TOP_GPA, _ = strconv.ParseFloat(gpa, 64)
	}
}

func GetStudentById(id int) (*models.Student, error) {
	return repository.FindById(id)
}

func GetAllStudents() []models.Student {
	return repository.FindAll()
}

func GetTopStudents() []models.Student {
	var topStudents []models.Student
	students := repository.FindAll()
	// all students where GPA is more than 4.5
	for _, s := range students {
		if s.GPA > TOP_GPA {
			topStudents = append(topStudents, s)
		}
	}
	return topStudents
}
