package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"student-backend/service"
)

func StartServer() {
	// registering the end-point/route (/greet) with the request multiplexer
	http.HandleFunc("/students", studentHandler)
	http.HandleFunc("/students/top", topStudentHandler)

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
		json.NewEncoder(w).Encode(service.GetAllStudents()) // returning slice
		return
	} else { // for single student
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Println("Error occoured while parsing String to int: ", err)
			http.Error(w, "Invalid Id", http.StatusBadRequest) // response for end user, in API response
			return
		}

		if student, err := service.GetStudentById(id); err != nil {
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

func topStudentHandler(w http.ResponseWriter, r *http.Request) {
	students := service.GetTopStudents()
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(students)
}
