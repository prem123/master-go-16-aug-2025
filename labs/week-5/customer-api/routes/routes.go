package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	r := mux.NewRouter()

	r.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{id}", getCustomer).Methods(http.MethodGet)

	log.Println("Starting server on 8080 port....")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func WriteResponse(w http.ResponseWriter, response any, httpStatusCode int) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}
