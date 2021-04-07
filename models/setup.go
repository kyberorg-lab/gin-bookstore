package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=db user=slon dbname=bookstore port=5432 password=slonopass sslmode=disable"
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to launch Database migrations!")
	}

	DB = database
}
