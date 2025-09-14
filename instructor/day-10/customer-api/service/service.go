package service

import (
	"customer-api/model"
)

type CustomerService struct {
	// repo repository.CustomerRepositoryDB
	repo CustomerRepository
}

func (cs *CustomerService) GetAllCustomer() ([]model.Customer, error) {

	// conversion should happen of status from 0 -> In-Active
	customers, err := cs.repo.FindAll()
	if err != nil {
		return nil, err
	}
	for idx, c := range customers {
		if c.Status == "0" {
			customers[idx].Status = "In-Active"
		}
		if c.Status == "1" {
			customers[idx].Status = "Active"
		}
	}
	return customers, nil
}

func (cs *CustomerService) GetCustomer(id string) (*model.Customer, error) {
	return cs.repo.FindById(id)
}

func (cs *CustomerService) AddCustomer(customer model.Customer) (int64, error) {
	return cs.repo.Save(customer)
}

// helper for service creation
func NewCustomerService(repo CustomerRepository) CustomerService {
	return CustomerService{repo}
}

type CustomerRepository interface {
	FindAll() ([]model.Customer, error)
	FindById(id string) (*model.Customer, error)
	Save(c model.Customer) (int64, error)
}
