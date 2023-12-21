package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global variable to later access the connections from main
var DB *gorm.DB

func DBconnetion() {
	// Load env variables from a .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connection string
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	// Open the connection to the database
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}

}
