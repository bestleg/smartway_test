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
	router.GET("/delete/:id", personRemove)
	router.GET("/personfromcompany/:id", showPersonsFromCompany)
	router.GET("/personfromdepartment/:name", showPersonsFromDepartment)
	router.POST("/", personAdd)
	router.PUT("/:id", personUpdate)

	http.ListenAndServe(":8000", router)
}
