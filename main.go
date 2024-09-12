package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/tnqbao/gau_services/config"
	"github.com/tnqbao/gau_services/routes"
)

func main() {
	db := config.InitDB()
	router := routes.SetupRouter(db)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Run(":8443")
}
