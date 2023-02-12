package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Product struct {
	gorm.Model
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
	Price       float64        `gorm:"not null" json:"price"`
	Discount    float64        `json:"discount"`
	Category    Category       `gorm:"foreignkey:CategoryID" json:"category"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	Sizes       pq.StringArray `gorm:"type:text[]" json:"sizes"`
	Colors      pq.StringArray `gorm:"type:text[]" json:"colors"`
	ImagesURL   pq.StringArray `gorm:"type:text[]" json:"images"`
}
