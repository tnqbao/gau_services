package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tnqbao/gau_services/models"
	"gorm.io/driver/mysql"
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
	mysql_user := getSecret("/run/secrets/mysql_user")
	mysql_password := getSecret("/run/secrets/mysql_password")
	mysql_host := getSecret("/run/secrets/mysql_host")
	database_name := getSecret("/run/secrets/db_name")

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
