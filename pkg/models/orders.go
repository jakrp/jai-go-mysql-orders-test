package models

import (
	"github.com/jakrp/jai-go-mysql-orders-test.git/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Order struct {
	gorm.Model
	Instrutment string `gorm:""json:"Instrutment"`
	Ordertype   string `json:"Ordertype"`
	Buysell     string `json:"Buysell"`
	Quantity    string `json:"Quantity"`
	Price       string `json:"Price"`
	Status      string `json:"Status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Order{})
}

func (b *Order) CreateOrder() *Order {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllOrders() []Order {
	var Orders []Order
	db.Find(&Orders)
	return Orders
}

func GetOrderById(Id int64) (*Order, *gorm.DB) {
	var getOrder Order
	db := db.Where("ID=?", Id).Find(&getOrder)
	return &getOrder, db
}

func DeleteOrder(ID int64) Order {
	var order Order
	db.Where("ID=?", ID).Delete(order)
	return order
}
