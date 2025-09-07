package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"student-backend/models"
	"student-backend/service"

	"github.com/gorilla/mux"
)

func StartServer() {
	// registering the end-point/route (/greet) with the request multiplexer

	// students
	// CRUD
	// Add 			-> /students + http method POST
	// updating -> /students + http method PUT (replace all) / PATCH (part)
	// get all  -> /students + http method GET
	// getBy id -> /students/{id} + http method GET -> /students/1 /students/2 /students/3
	// delete   -> /students + http method DELETE
	// top students -> /students/top

	r := mux.NewRouter()

	r.HandleFunc("/students", studentsHandler).Methods(http.MethodGet)
	r.HandleFunc("/students/{id:[0-9]+}", studentHandler).Methods(http.MethodGet)

	r.HandleFunc("/students", newStudentHandler).Methods(http.MethodPost)

	// HomeWork : Implement this method
	// r.HandleFunc("/students/{id}", updateStudentHandler).Methods(http.MethodPut)
	// r.HandleFunc("/students/{id}", deleteStudentHandler).Methods(http.MethodDelete)

	r.HandleFunc("/students/top", topStudentHandler).Methods(http.MethodGet)

	log.Println("Starting server on 8080 port....")
	http.ListenAndServe("localhost:8080", r)
}

// REST -> Representational State Transfer Protocol
func studentHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	studentId := v["id"]
	id, _ := strconv.Atoi(studentId)

	if student, err := service.GetStudentById(id); err != nil {
		WriteResponse(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
	} else {
		WriteResponse(w, student, http.StatusOK)
	}
}

func newStudentHandler(w http.ResponseWriter, r *http.Request) {
	var newStudent models.Student

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() //strict
	err := decoder.Decode(&newStudent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// here we will save the record
	id := service.AddNewStudent(newStudent)

	WriteResponse(w, map[string]int{"id": id}, http.StatusCreated)
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
func studentsHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id") // this is optional

	if idString == "" { // for all students
		WriteResponse(w, service.GetAllStudents(), http.StatusOK)
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
			WriteResponse(w, student, http.StatusOK)
		}
	}
}

func topStudentHandler(w http.ResponseWriter, r *http.Request) {
	students := service.GetTopStudents()
	WriteResponse(w, students, http.StatusOK)
}

func WriteResponse(w http.ResponseWriter, response any, httpStatusCode int) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}
