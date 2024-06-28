package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

// InitDB initializes the database connection
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var connErr error
	dbURL := os.Getenv("DATABASE_URL")
	DB, connErr = pgxpool.Connect(context.Background(), dbURL)
	if connErr != nil {
		log.Fatal("Error connecting to database: ", connErr)
	}

	fmt.Println("Connected to database")
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
