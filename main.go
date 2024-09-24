package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tnqbao/gau_services/config"
	"github.com/tnqbao/gau_services/routes"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found!")
	}
	db := config.InitDB()
	router := routes.SetupRouter(db)
	router.Run(":8080")
}
