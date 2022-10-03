package routes

import (
	"github.com/gorilla/mux"
	"github.com/jakrp/jai-go-mysql-orders-test.git/pkg/controllers"
)

var RegisterOrderRoutes = func(router *mux.Router) {
	router.HandleFunc("/orders/", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/orders/{orderid}", controllers.GetOrderById).Methods("GET")
	router.HandleFunc("/orders/{orderid}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderid}", controllers.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/orders/trigger/{orderid}", controllers.TriggerOrder).Methods("POST")
}
