package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"transactions/config"
	"transactions/models"
)

var DB *gorm.DB

func InitToDatabase() *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(config.Configure.DB), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB, err: ", err.Error())
	}

	err = DB.AutoMigrate(&models.User{}, &models.Transaction{})
	if err != nil {
		log.Println("cannot migrate models to db, err: ", err.Error())
	}
	return DB
}
