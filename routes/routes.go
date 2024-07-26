package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_services/controllers"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/", controllers.CreateUser)
	}

	return r
}
