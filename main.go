package main

import (
	"github.com/gin-contrib/cors"
	"github.com/tnqbao/gau_services/config"
	"github.com/tnqbao/gau_services/routes"
)

func main() {

	db := config.InitDB()

	router := routes.SetupRouter(db)

	router.Use(cors.Default())

	router.Run(":8080")
}
