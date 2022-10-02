package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jakrp/jai-go-mysql-orders-test.git/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterOrderRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
