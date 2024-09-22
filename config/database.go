package config

import (
	"fmt"
	"log"
	"os"

	"github.com/tnqbao/gau_services/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is missing", key)
	}
	return value
}

func InitDB() *gorm.DB {
	mysql_user := getEnv("MYSQL_USER")
	mysql_password := getEnv("MYSQL_PASSWORD")
	mysql_host := getEnv("MYSQL_HOST")
	database_name := getEnv("MYSQL_DATABASE")

	if mysql_user == "" || mysql_password == "" || mysql_host == "" || database_name == "" {
		log.Fatal("One or more required secrets are missing")
	}

	fmt.Printf("DB connect status: %s:%s@tcp(%s:3306)/%s\n", mysql_user, mysql_password, mysql_host, database_name)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysql_user, mysql_password, mysql_host, database_name)

	var err error

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
