package database

import (
	"Golang-fiber/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("DB_TYPE") == "mysql" {
		dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT_MYSQL") + ")/" + os.Getenv("DB_NAME")
		connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect mysql database")
		}
		DB = connection
		err = connection.AutoMigrate(&models.User{}, &models.Blog{})
		if err != nil {
			return
		}
	}

	if os.Getenv("DB_TYPE") == "postgres" {
		dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_HOST") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT_POSTGRES") + " sslmode=disable TimeZone=Asia/Jakarta"
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		DB = connection
		err = connection.AutoMigrate(&models.User{}, &models.Blog{})
		if err != nil {
			return
		}
	}
}
