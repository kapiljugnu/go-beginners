package domain

import "errors"

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}

var CustomerNotFound error = errors.New("customer not found")
var NoCustomerExists error = errors.New("no customer exists")
var CustomerIdAlreadyExist error = errors.New("customer id already exist")
