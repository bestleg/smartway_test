package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"smartway_test/models"
)

var db *gorm.DB

func main() {
	dsn := "host=localhost user=admin password=123 dbname=smartway port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Person{}, &models.Department{}, &models.Passport{})

	router := httprouter.New()
	router.GET("/", GetAll)
	router.GET("/get/:id", GetByID)
	router.GET("/delete/:id", RemovePerson)
	router.GET("/personfromcompany/:id", ShowPersonsFromCompany)
	router.GET("/personfromdepartment/:name", ShowPersonsFromDepartment)
	router.POST("/", TestPost)

	http.ListenAndServe(":8000", router)
}
