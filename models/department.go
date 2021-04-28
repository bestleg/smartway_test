package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name  string `json:"name" gorm:"name"`
	Phone string `json:"phone" gorm:"phone"`
}
