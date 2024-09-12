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

	apiRoutes := r.Group("/api")
	{
		userRoutes := apiRoutes.Group("/users")
		{
			userRoutes.POST("/register", controllers.Register)
			userRoutes.GET("/:id", controllers.GetUserById)
			userRoutes.DELETE("/:id", controllers.DeleteUserById)
			userRoutes.PUT("/update/:id", controllers.UpdateUserInformation)
			userRoutes.POST("/login", controllers.Authentication)

		}
	}

	return r
}
