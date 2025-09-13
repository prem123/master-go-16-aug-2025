package service

import (
	"customer-api/model"
	"customer-api/repository"
)

type CustomerService struct {
	repo repository.CustomerRepositoryDB
}

func (cs *CustomerService) GetAllCustomer() ([]model.Customer, error) {
	return cs.repo.FindAll()
}

func (cs *CustomerService) GetCustomer(id string) (*model.Customer, error) {
	repo := repository.NewCustomerRepositoryStub()
	return repo.FindById(id)
}

// helper for service creation

func NewCustomerService(repo repository.CustomerRepositoryDB) CustomerService {
	return CustomerService{repo}
}
