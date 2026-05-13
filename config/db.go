package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to load .env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Database: %s", err)
	}
	fmt.Println("Database connection successful.")

	return db, nil
}
