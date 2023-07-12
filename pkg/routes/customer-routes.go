package routes

import (
	"customerapi/pkg/controllers"

	"github.com/gorilla/mux"
)

var CustomerInfoRouter = func(router *mux.Router) {
	router.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{customerId}", controllers.GetCustomer).Methods("GET")
	router.HandleFunc("/customer/create", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/update/{customerId}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customer/delete/{customerId}", controllers.DeleteCustomer).Methods("DELETE")
}
