package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID    int       `json:"user_id"`
	Products  []Product `json:"products" gorm:"many2many:order_products;"`
	TotalCost float64   `json:"total_cost"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
}
