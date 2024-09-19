package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/tnqbao/gau_services/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getSecret(path string) string {
	secret, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error opening secret file: %v", err)
	}
	return strings.TrimSpace(string(secret))
}

func InitDB() *gorm.DB {
	username_db := getSecret("/run/secrets/db_username")
	password_db := getSecret("/run/secrets/mysql_root_password")
	address_db := getSecret("/run/secrets/db_address")
	database_name := getSecret("/run/secrets/db_name")

	if username_db == "" || password_db == "" || address_db == "" || database_name == "" {
		log.Fatal("One or more required secrets are missing")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username_db, password_db, address_db, database_name)
	fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username_db, password_db, address_db, database_name)

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
