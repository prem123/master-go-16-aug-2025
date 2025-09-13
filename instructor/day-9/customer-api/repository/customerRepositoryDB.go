package repository

import (
	"customer-api/model"
	"database/sql"
	"log"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (cr CustomerRepositoryDB) FindAll() ([]model.Customer, error) {

	// this is an expensive operation
	// db := NewDBConnection() // this is for now, and is wrong design

	q := "select customer_id, name, date_of_birth, city, zipcode, status from customers"

	rows, err := cr.client.Query(q)
	if err != nil {
		log.Println("Error occoured while fetching customers from DB. ", err.Error())
		return nil, err
	}

	customers := make([]model.Customer, 0)

	for rows.Next() {
		var c model.Customer
		err := rows.Scan(&c.Id, &c.Name, &c.DateofBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error occured while scaning customer to rows. ", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// this is a helper function and is helping us in creating an instance of
// CustomerRepositoryDB. This also exposes that to instantiate repo we also
// need the client
func NewCustomerRepositoryDB(db *sql.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{client: db}
}
