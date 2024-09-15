package main

import (
	"github.com/tnqbao/gau_services/config"
	"github.com/tnqbao/gau_services/routes"
)

func main() {
	db := config.InitDB()
	router := routes.SetupRouter(db)
	router.Run(":8080")
}
