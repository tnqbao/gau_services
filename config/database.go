package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username_db := os.Getenv("USERNAME_DATABASE")
	password_db := os.Getenv("PASSWORD_DATABASE")
	address_db := os.Getenv("ADDRESS_DATABASE")
	database_name := os.Getenv("DATABSE_NAME")

	if username_db == "" || password_db == "" || address_db == "" || database_name == "" {
		log.Fatal("One or more required environment variables are missing")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username_db, password_db, address_db, database_name)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected")

	err = DB.AutoMigrate(&models.User{}, &models.UserInformation{}, &models.UserAuthentication{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return DB
}
