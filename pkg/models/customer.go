package models

import (
	"customerapi/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Customer struct {
	gorm.Model
	FirstName   string `gorm:""json:"firstname"`
	LastName    string `json:"lastname"`
	Contact     string `json:"contact"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	CreditLimit int    `json:"creditlimit"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Customer{})
}

func (c *Customer) CreateCustomer() *Customer {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetCustomers() []Customer {
	var Customers []Customer
	db.Find(&Customers)
	return Customers
}

func GetCustomer(Id int64) (*Customer, *gorm.DB) {
	var GetCustomer Customer
	db.Where("ID=?", Id).Find(&GetCustomer)
	return &GetCustomer, db
}

func DeleteCustomer(Id int64) Customer {
	var Customer Customer
	db.Where("ID=?", Id).Delete(Customer)
	return Customer
}
