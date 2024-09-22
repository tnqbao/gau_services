package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tnqbao/gau_services/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getSecret(path string) string {
	secret, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error opening secret file: %v", err)
	}
	return strings.TrimSpace(string(secret))
}

func InitDB() *gorm.DB {
	pg_user := getSecret("/run/secrets/pg_user")
	pg_password := getSecret("/run/secrets/pg_password")
	pg_host := "postgres"
	database_name := "gau_services_db"

	if pg_user == "" || pg_password == "" || pg_host == "" || database_name == "" {
		log.Fatal("One or more required secrets are missing")
	}

	fmt.Printf("DB connect status: %s:%s@tcp(%s:5432)/%s\n", pg_user, pg_password, pg_host, database_name)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh", pg_host, pg_user, pg_password, database_name)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
