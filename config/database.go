package config

import (
	"log"

	"github.com/tnqbao/gau_services/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "jgdnwaaqhosting_tnqbao:Gau_12345@tcp(202.92.4.30:3306)/jgdnwaaqhosting_gau_service"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected")
	DB.AutoMigrate(&models.User{}, &models.UserInformation{}, &models.UserAuthentication{})
	return DB
}
