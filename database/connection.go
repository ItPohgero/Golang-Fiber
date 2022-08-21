package database

import (
	"Golang-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root@/golang_fiber"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
