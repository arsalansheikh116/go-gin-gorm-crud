package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-crud/models"
)

var DB *gorm.DB

func InitializeDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("HOST"), GetEnv("PORT"), GetEnv("USER"), GetEnv("PASSWORD"), GetEnv("DBNAME"))
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Automigrate models
	err = DB.AutoMigrate(models.GetModels()...)
	if err != nil {
		log.Fatal("Failed to automigrate database:", err)
	}
}
