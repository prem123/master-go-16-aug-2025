package service

import (
	"customer-api/model"
	"customer-api/repository"
)

func GetAllCustomer() ([]model.Customer, error) {
	repo := repository.NewCustomerRepositoryStub()
	return repo.FindAll()
}

func GetCustomer(id string) (*model.Customer, error) {
	repo := repository.NewCustomerRepositoryStub()
	return repo.FindById(id)
}
