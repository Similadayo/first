package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string  `gorm:"type:varchar(100);not null"`
	LastName  string  `gorm:"type:varchar(100);not null"`
	UserName  string  `gorm:"type:varchar(100);not null"`
	Email     string  `gorm:"type:varchar(100);not null;unique"`
	Password  string  `gorm:"type:varchar(255);not null"`
	Phone     string  `gorm:"type:varchar(100);not null"`
	Address   string  `gorm:"type:varchar(100);not null"`
	Role      string  `gorm:"type:varchar(100);not null"`
	Orders    []Order `gorm:"ForeignKey:UserID"`
}
