package controller

import (
	"errors"
	"fmt"

	"customerapp/domain"
)

type CustomerController struct {
	Store domain.CustomerStore
}

func LogError(err error) {
	fmt.Println("Error:", err)
}

func (c CustomerController) Add(cust domain.Customer) {
	err := c.Store.Create(cust)
	if errors.Is(err, domain.CustomerIdAlreadyExist) {
		fmt.Println("CustomerIdAlreadyExist", err)
	}
	if err != nil {
		LogError(err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (c CustomerController) Change(id string, cust domain.Customer) {
	err := c.Store.Update(id, cust)
	if err != nil {
		LogError(err)
		return
	}
	fmt.Println("Customer with id", id, "updated successfully")
}

func (c CustomerController) FindById(id string) {
	customer, err := c.Store.GetById(id)
	if err != nil {
		LogError(err)
		return
	}
	fmt.Println("details of customer with id", id, "are", customer)
}

func (c CustomerController) GetAll() {
	customers, err := c.Store.GetAll()
	if err != nil {
		LogError(err)
		return
	}
	for _, v := range customers {
		fmt.Println(v)
	}
}

func (c CustomerController) Remove(id string) {
	err := c.Store.Delete(id)
	if err != nil {
		LogError(err)
		return
	}
	fmt.Println("Customer with id", id, "removed successfully")
}
