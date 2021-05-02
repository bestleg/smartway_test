package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name         string     `json:"first_name" gorm:"first_name"`
	Surname      string     `json:"last_name" gorm:"last_name"`
	Phone        int        `json:"phone" gorm:"phone"`
	CompanyID    int        `json:"company_id" gorm:"company_id"`
	DepartmentID int        `json:"department_id"`
	Department   Department `json:"department"`
	PassportID   int        `json:"passport_id"`
	Passport     Passport   `json:"passport"`
}
