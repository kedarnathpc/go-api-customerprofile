package main

import (
	// "Customer Profiles APIs/pkg/routes"
	"customerapi/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.CustomerInfoRouter(r)
	http.Handle("/", r)
	fmt.Println("Starting server at port :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
