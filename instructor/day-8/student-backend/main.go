package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"student-backend/models"
)

var sm = models.StudentManager{} // Create a new StudentManager instance

func main() {
	sm.Enroll(models.Student{Id: 1, Name: "Alice", Age: 20, GPA: 3.5, Password: "pass1"})
	sm.Enroll(models.Student{Id: 2, Name: "Bob", Age: 21, GPA: 3.2})
	sm.Enroll(models.Student{Id: 3, Name: "Charlie", Age: 22, GPA: 3.9})

	// registering the end-point/route (/greet) with the request multiplexer
	http.HandleFunc("/students", studentHandler)

	log.Println("Starting server on 8080 port....")
	http.ListenAndServe("localhost:8080", nil)
}

// the responsibility of handler
// 1. to receive request
// 2. validate incoming data
// 3. delegate request to another component for processing (heavy lifting)
// 4. Return response

/*
return single student incase of ID and all when no ID is present
/students?id=all -> 400 Bad request
/students?id=10 -> 404 Not Found
*/
func studentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idString := r.URL.Query().Get("id") // this is optional

	if idString == "" { // for all students
		json.NewEncoder(w).Encode(sm.Students) // returning slice
		return
	} else { // for single student
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Println("Error occoured while parsing String to int: ", err)
			http.Error(w, "Invalid Id", http.StatusBadRequest) // response for end user, in API response
			return
		}

		if student, err := getStudentById(id); err != nil {
			// e := map[string]string{
			// 	"error_code": "10202", // internal application error code
			// 	"message":    err.Error(),
			// }
			// w.WriteHeader(http.StatusBadRequest)
			// json.NewEncoder(w).Encode(e)
			http.Error(w, err.Error(), http.StatusNotFound) // response for end user, in API response
		} else {
			json.NewEncoder(w).Encode(student) // single object
		}
	}
}

func getStudentById(id int) (*models.Student, error) {
	for _, student := range sm.Students {
		if student.Id == id {
			return &student, nil
		}
	}
	return nil, errors.New("id not found")
}
