package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name         string `json:"first_name" gorm:"first_name"`
	Surname      string `json:"last_name" gorm:"last_name"`
	Phone        string `json:"phone" gorm:"phone"`
	DepartmentID int
	Department   Department
	PassportID   int
	Passport     Passport
}
