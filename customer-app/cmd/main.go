package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"customerapp/controller"
	"customerapp/router"
	"customerapp/store"
)

func main() {
	customerController := controller.CustomerController{
		Store: store.NewMapStore(),
	}

	r := mux.NewRouter()
	router.IntiailizeCustomerRouter(r, customerController)

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", r)

}
