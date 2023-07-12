package controllers

import (
	"customerapi/pkg/models"
	"customerapi/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewCustomer models.Customer

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers := models.GetCustomers()

	res, _ := json.Marshal(customers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]

	ID, err := strconv.ParseInt(customerId, 0, 0)
	checkError(err)

	customerDetails, _ := models.GetCustomer(ID)

	res, _ := json.Marshal(customerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	CreateCustomer := &models.Customer{}
	utils.ParseBody(r, CreateCustomer)
	c := CreateCustomer.CreateCustomer()

	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	ID, err := strconv.ParseInt(customerId, 0, 0)
	checkError(err)

	customer := models.DeleteCustomer(ID)

	res, _ := json.Marshal(customer)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var updatecustomer = &models.Customer{}
	utils.ParseBody(r, updatecustomer)

	vars := mux.Vars(r)
	customerId := vars["customerId"]
	ID, err := strconv.ParseInt(customerId, 0, 0)
	checkError(err)

	customerDetails, db := models.GetCustomer(ID)

	if updatecustomer.FirstName != "" {
		customerDetails.FirstName = updatecustomer.FirstName
	}

	if updatecustomer.LastName != "" {
		customerDetails.LastName = updatecustomer.LastName
	}

	if updatecustomer.Contact != "" {
		customerDetails.Contact = updatecustomer.Contact
	}

	if updatecustomer.Address != "" {
		customerDetails.Address = updatecustomer.Address
	}

	if updatecustomer.City != "" {
		customerDetails.City = updatecustomer.City
	}

	if updatecustomer.State != "" {
		customerDetails.State = updatecustomer.State
	}

	if updatecustomer.Country != "" {
		customerDetails.Country = updatecustomer.Country
	}

	if updatecustomer.CreditLimit != 0 {
		customerDetails.CreditLimit = updatecustomer.CreditLimit
	}

	db.Save(&customerDetails)

	res, _ := json.Marshal(customerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
