package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"customerapp/domain"
)

type CustomerController struct {
	Store domain.CustomerStore
}

func (c CustomerController) Create(w http.ResponseWriter, r *http.Request) {
	var cust domain.Customer
	if err := json.NewDecoder(r.Body).Decode(&cust); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := c.Store.Create(cust)
	if errors.Is(err, domain.ErrCustomerIdAlreadyExist) {
		fmt.Println("CustomerIdAlreadyExist", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		LogError(err)
		return
	}
	jsonResponse(w, http.StatusCreated, "customer created successfuly")
}

func (c CustomerController) Change(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var cust domain.Customer
	if err := json.NewDecoder(r.Body).Decode(&cust); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := c.Store.Update(id, cust)
	if err != nil {
		LogError(err)
		if errors.Is(err, domain.ErrCustomerNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, http.StatusNoContent, "Customer with id "+id+" updated successfully")
}

func (c CustomerController) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	customer, err := c.Store.GetById(id)
	if err != nil {
		LogError(err)
		if errors.Is(err, domain.ErrCustomerNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, http.StatusOK, customer)
}

func (c CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := c.Store.GetAll()
	if err != nil {
		LogError(err)
		if errors.Is(err, domain.ErrNoCustomerExists) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, http.StatusOK, customers)
}

func (c CustomerController) Remove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := c.Store.Delete(id)
	if err != nil {
		LogError(err)
		if errors.Is(err, domain.ErrCustomerNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, http.StatusOK, "Customer with id "+id+" removed successfully")
}
