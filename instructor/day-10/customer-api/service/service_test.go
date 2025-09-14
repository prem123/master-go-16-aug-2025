package service

import (
	"customer-api/model"
	"testing"
)

// interaction test -> using mocks
// mocks can only be created from interface
func Test_get_all_customers_should_call_repo_FindAll(t *testing.T) {
	// AAA

	// Arrange
	mockRepo := NewMockCustomerRepository(t)
	svc := NewCustomerService(mockRepo)
	expected := []model.Customer{{Id: "1", Name: "testmock"}}

	// Assert -> you define your interaction expectation and assertion together
	mockRepo.EXPECT().FindAll().Return(expected, nil)

	// ACT
	svc.GetAllCustomer()

	// You dont have to test the response from GetAllCustomers because it will always be same
	// and the test will be an invalid test

	// customers, err := svc.GetAllCustomer()
	// fmt.Println(customers)
	// fmt.Println(err)
}

func Test_get_all_customers_should_return_valid_status(t *testing.T) {
	// Arrange
	mockRepo := NewMockCustomerRepository(t)
	svc := NewCustomerService(mockRepo)
	repoCustomers := []model.Customer{
		{Id: "1", Name: "testmock1", Status: "0"},
		{Id: "2", Name: "testmock2", Status: "1"},
	}

	// Assert -> you define your interaction expectation and assertion together
	mockRepo.EXPECT().FindAll().Return(repoCustomers, nil)

	// ACT
	actual, _ := svc.GetAllCustomer()

	expected := []model.Customer{
		{Id: "1", Name: "testmock1", Status: "In-Active"},
		{Id: "2", Name: "testmock2", Status: "Active"},
	}

	for idx := range expected {
		if expected[idx].Status != actual[idx].Status {
			t.Error("Status not matching")
		}
	}

	// if expected[0].Status != actual[0].Status && expected[1].Status != actual[1].Status {
	// t.Error("Status not matching")
	// }
}
