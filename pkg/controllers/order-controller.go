package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jakrp/jai-go-mysql-orders-test.git/pkg/models"
	"github.com/jakrp/jai-go-mysql-orders-test.git/pkg/utils"
)

var NewOrder models.Order

func GetOrders(w http.ResponseWriter, r *http.Request) {
	go GetOrderMultiple(w, r)
	go GetOrderMultiple(w, r)
	go GetOrderMultiple(w, r)
	GetOrderMultiple(w, r)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	go GetOrderByIdMultiple(w, r)
	go GetOrderByIdMultiple(w, r)
	go GetOrderByIdMultiple(w, r)
	GetOrderByIdMultiple(w, r)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	go CreateOrderMultiple(w, r)
	go CreateOrderMultiple(w, r)
	go CreateOrderMultiple(w, r)
	CreateOrderMultiple(w, r)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	go DeleteOrderMultiple(w, r)
	go DeleteOrderMultiple(w, r)
	go DeleteOrderMultiple(w, r)
	DeleteOrderMultiple(w, r)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	go UpdateOrderMultiple(w, r)
	go UpdateOrderMultiple(w, r)
	go UpdateOrderMultiple(w, r)
	UpdateOrderMultiple(w, r)
}

func GetOrderMultiple(w http.ResponseWriter, r *http.Request) {
	newOrders := models.GetAllOrders()
	res, _ := json.Marshal(newOrders)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetOrderByIdMultiple(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	ID, err := strconv.ParseInt(orderId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	orderDetails, _ := models.GetOrderById(ID)
	res, _ := json.Marshal(orderDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateOrderMultiple(w http.ResponseWriter, r *http.Request) {
	CreateOrder := &models.Order{}
	utils.ParseBody(r, CreateOrder)
	b := CreateOrder.CreateOrder()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteOrderMultiple(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	ID, err := strconv.ParseInt(orderId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	order := models.DeleteOrder(ID)
	res, _ := json.Marshal(order)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateOrderMultiple(w http.ResponseWriter, r *http.Request) {
	var updateOrder = &models.Order{}
	utils.ParseBody(r, updateOrder)
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	ID, err := strconv.ParseInt(orderId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	orderDetails, db := models.GetOrderById(ID)

	if updateOrder.Instrutment != "" {
		orderDetails.Instrutment = updateOrder.Instrutment
	}

	if updateOrder.Ordertype != "" {
		orderDetails.Ordertype = updateOrder.Ordertype
	}

	if updateOrder.Buysell != "" {
		orderDetails.Buysell = updateOrder.Buysell
	}

	if updateOrder.Quantity != "" {
		orderDetails.Quantity = updateOrder.Quantity
	}

	if updateOrder.Price != "" {
		orderDetails.Price = updateOrder.Price
	}

	db.Save(&orderDetails)
	res, _ := json.Marshal(orderDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
