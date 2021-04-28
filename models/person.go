package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name         string `json:"first_name" gorm:"first_name"`
	Surname      string `json:"last_name" gorm:"last_name"`
	Phone        string `json:"phone" gorm:"phone"`
	CompanyID    int    `json:"company_id" gorm:"company_id"`
	DepartmentID int
	Department   Department
	PassportID   int
	Passport     Passport
}
