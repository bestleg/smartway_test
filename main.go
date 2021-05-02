package main

import (
	"net/http"
	"smartway_test/database"
	"smartway_test/routers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	database.MigrateDB()

	router := httprouter.New()
	router.GET("/delete/:id", routers.PersonRemove)
	router.GET("/personfromcompany/:id", routers.ShowPersonsFromCompany)
	router.GET("/personfromdepartment/:name", routers.ShowPersonsFromDepartment)
	router.POST("/", routers.PersonAdd)
	router.PUT("/:id", routers.PersonUpdate)

	http.ListenAndServe(":8000", router)
}
