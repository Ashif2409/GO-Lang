package database

import (
	"fmt"
	"log"
	"os"

	models "example.com1/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load database credentials from environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "test_db"),
		getEnv("DB_PORT", "5432"),
	)

	// Connect to PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate models
	db.AutoMigrate(&models.Book{}) // Add your models here

	DB = db
	log.Println("Connected to database successfully")
}

// Helper function to get environment variables with default values
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
