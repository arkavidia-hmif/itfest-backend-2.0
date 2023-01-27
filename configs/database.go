package configs

import (
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"itfest-backend-2.0/models"
)

type Database struct {
	connection *gorm.DB
	once       sync.Once
}

func (database *Database) lazyInit() {
	database.once.Do(func() {
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

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
		if err != nil {
			panic("Cannot connect database")
		}

		db.AutoMigrate(
			&models.Profile{},
			&models.Merchandise{},
			&models.Log{},
			&models.Game{},
			&models.Clue{},
		)

		db.AutoMigrate(
			&models.User{},
		)
		// Assign To Struct
		database.connection = db
	})
}

func (database *Database) GetConnection() *gorm.DB {
	database.lazyInit()
	return database.connection
}

var DB = &Database{}
