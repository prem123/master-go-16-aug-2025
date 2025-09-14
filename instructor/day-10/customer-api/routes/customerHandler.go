package routes

import (
	"customer-api/model"
	"customer-api/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllCustomer()

	if err != nil {
		WriteResponse(w, err.Error(), http.StatusInternalServerError)
	} else {
		WriteResponse(w, customers, http.StatusOK)
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		WriteResponse(w, customer, http.StatusOK)
	}
}

func (ch *CustomerHandler) addNewCustomer(w http.ResponseWriter, r *http.Request) {

	var customer model.Customer
	// handle error as homework
	json.NewDecoder(r.Body).Decode(&customer)

	newCustomerId, err := ch.service.AddCustomer(customer)
	if err != nil {
		WriteResponse(w, err, http.StatusInternalServerError)
	} else {
		WriteResponse(w, map[string]int64{"id": newCustomerId}, http.StatusCreated)
	}
}

// helper
func NewCustomerHandler(service service.CustomerService) CustomerHandler {
	return CustomerHandler{service}
}
