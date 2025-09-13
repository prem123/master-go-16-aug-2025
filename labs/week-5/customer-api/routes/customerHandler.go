package routes

import (
	"customer-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := service.GetAllCustomer()

	if err != nil {
		WriteResponse(w, err.Error(), http.StatusInternalServerError)
	} else {
		WriteResponse(w, customers, http.StatusOK)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		WriteResponse(w, customer, http.StatusOK)
	}
}
