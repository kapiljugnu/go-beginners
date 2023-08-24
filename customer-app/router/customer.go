package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"customerapp/controller"
)

func IntiailizeCustomerRouter(r *mux.Router, c controller.CustomerController) {
	r.HandleFunc("/api/customers", c.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/customers/{id}", c.FindById).Methods(http.MethodGet)
	r.HandleFunc("/api/customers", c.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/customers/{id}", c.Remove).Methods(http.MethodDelete)
	r.HandleFunc("/api/customers/{id}", c.Change).Methods(http.MethodPatch)
}
