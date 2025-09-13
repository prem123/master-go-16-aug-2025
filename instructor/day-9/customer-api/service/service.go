package service

import (
	"customer-api/model"
)

type CustomerService struct {
	// repo repository.CustomerRepositoryDB
	repo CustomerRepository
}

func (cs *CustomerService) GetAllCustomer() ([]model.Customer, error) {
	return cs.repo.FindAll()
}

func (cs *CustomerService) GetCustomer(id string) (*model.Customer, error) {
	return cs.repo.FindById(id)
}

// helper for service creation
func NewCustomerService(repo CustomerRepository) CustomerService {
	return CustomerService{repo}
}

type CustomerRepository interface {
	FindAll() ([]model.Customer, error)
	FindById(id string) (*model.Customer, error)
}
