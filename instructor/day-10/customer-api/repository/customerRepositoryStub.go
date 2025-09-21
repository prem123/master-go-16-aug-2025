package repository

import (
	"customer-api/model"
	"errors"
)

type CustomerRepositoryStub struct {
	customers []model.Customer
}

func (s CustomerRepositoryStub) FindAll() ([]model.Customer, error) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) FindById(id string) (*model.Customer, error) {
	for _, c := range s.customers {
		if c.Id == id {
			return &c, nil
		}
	}
	return nil, errors.New("not found")
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []model.Customer{
		{Id: "11", Name: "Balraj", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "12", Name: "Birender", City: "Mumbai", Zipcode: "400001", DateofBirth: "2000-08-05", Status: "0"},
		{Id: "13", Name: "Kushal", City: "Chennai", Zipcode: "600001", DateofBirth: "2000-02-23", Status: "1"},
		{Id: "14", Name: "Vikram", City: "kolkata", Zipcode: "700064", DateofBirth: "2000-06-16", Status: "0"},
	}
	return CustomerRepositoryStub{customers}
}
