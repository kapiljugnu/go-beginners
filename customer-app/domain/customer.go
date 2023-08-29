package domain

import "errors"

//go:generate mockgen -destination=../mocks/mock_customer.go -package=mocks . CustomerStore

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

var ErrCustomerNotFound error = errors.New("customer not found")
var ErrNoCustomerExists error = errors.New("no customer exists")
var ErrCustomerIdAlreadyExist error = errors.New("customer id already exist")
