package main

import (
	"customerapp/controller"
	"customerapp/domain"
	"customerapp/store"
)

func main() {
	controller := controller.CustomerController{
		Store: store.NewMapStore(),
	}

	customer := domain.Customer{
		ID:    "cust101",
		Name:  "John",
		Email: "john.doe@fake.com",
	}

	duplicateCustomer := customer

	controller.Add(customer)
	controller.Add(duplicateCustomer)

	duplicateCustomer.Name = "duplicate john"

	controller.Change(customer.ID, duplicateCustomer)

	controller.Change("unknown id", duplicateCustomer)

	controller.FindById(customer.ID)

	anotherCustomer := domain.Customer{
		ID:    "cust102",
		Name:  "Jane",
		Email: "jane.doe@fake.com",
	}

	controller.Add(anotherCustomer)

	controller.GetAll()

	controller.Remove(customer.ID)
	controller.GetAll()

}
