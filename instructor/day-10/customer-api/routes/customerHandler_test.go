package routes

import (
	"customer-api/model"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_get_all_customers_should_return_200_OK(t *testing.T) {

	// Arrange
	mockSvc := NewMockCustomerService(t)
	ch := NewCustomerHandler(mockSvc)
	customers := []model.Customer{{Id: "testid1", Name: "test handler 1"}}
	mockSvc.EXPECT().GetAllCustomer().Return(customers, nil)

	r := mux.NewRouter()
	r.HandleFunc("/customers", ch.getAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	recorder := httptest.NewRecorder()

	// ACT
	r.ServeHTTP(recorder, request)

	// Assert
	if recorder.Result().StatusCode != http.StatusOK {
		t.Error("failed, because status code is not 200 OK...")
	}
}

func Test_get_all_customers_should_return_500_InternalServerError_in_case_of_error(t *testing.T) {

	// Arrange
	mockSvc := NewMockCustomerService(t)
	ch := NewCustomerHandler(mockSvc)
	mockSvc.EXPECT().GetAllCustomer().Return(nil, errors.New("some error"))

	r := mux.NewRouter()
	r.HandleFunc("/customers", ch.getAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	recorder := httptest.NewRecorder()

	// ACT
	r.ServeHTTP(recorder, request)

	// Assert
	if recorder.Result().StatusCode != http.StatusInternalServerError {
		t.Error("failed, because status code is not 500 InternalServerError...")
	}
}
