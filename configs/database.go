package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"itfest-backend-2.0/models"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DEV_HOST")
	port := os.Getenv("DEV_PORT")
	dbname := os.Getenv("DEV_DBNAME")
	username := os.Getenv("DEV_USERNAME")
	password := os.Getenv("DEV_PASSWORD")

	dsn := "host=" + host
	dsn += " user=" + username
	dsn += " password=" + password
	dsn += " dbname=" + dbname
	dsn += " port=" + port
	dsn += " sslmode=disable"

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}

	DB.AutoMigrate(&models.Account{}, &models.User{}, &models.Log{}, &models.Profile{})

	fmt.Println("Database connected")
}
