package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	datasourcename := os.Getenv("DB_URL")
	if datasourcename == "" {
		panic("DB_URL not set in environment variables")
	}

	database, err := gorm.Open(postgres.Open(datasourcename), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	fmt.Println("Database connected")
}
