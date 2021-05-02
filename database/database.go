package database

import (
	"smartway_test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func MigrateDB() {
	dsn := "host=localhost user=admin password=123 dbname=smartway port=5432 sslmode=disable"
	var err error
	DataBase, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DataBase.AutoMigrate(&models.Person{}, &models.Department{}, &models.Passport{})
}
