package db

import (
	"authorisation_app/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := DB.AutoMigrate(&models.TUser{}, &models.TProduct{}); err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	if DB == nil {
		return
	}

	sqlDB, err := DB.DB()

	if err != nil {
		log.Println("Error getting underlying SQL DB:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Error closing the database:", err)
	}
}
