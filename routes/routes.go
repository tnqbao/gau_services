package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_services/controllers"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	apiRoutes := r.Group("/api")
	{
		userRoutes := apiRoutes.Group("/users")
		{

			userRoutes.GET("/:id", controllers.GetUserById)
			userRoutes.DELETE("/:id", controllers.DeleteUserById)
			userRoutes.PUT("/update/:id", controllers.UpdateUserInformation)

		}

		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/register", controllers.Register)
			authRoutes.POST("/login", controllers.Authentication)
		}
	}

	return r
}
