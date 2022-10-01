package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Order struct {
	gorm.Model
	Instrutment string  `gorm:""json:"instrutment"`
	Ordertype   string  `json:"ordertype"`
	Buysell     string  `json:"buysell"`
	Quantity    float32 `json:"quantity"`
	Price       float32 `json:"price"`
}
