package routes

import (
	"github.com/gin-gonic/gin"
	api_authed_user "github.com/tnqbao/gau_services/api/authed/user"
	api_public_auth "github.com/tnqbao/gau_services/api/public/auth"
	"github.com/tnqbao/gau_services/middlewares"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	apiRoutes := r.Group("/api")
	{
		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.Use(middlewares.AuthMiddleware())
			userRoutes.GET("/:id", api_authed_user.GetUserById)
			userRoutes.DELETE("/:id", api_authed_user.DeleteUserById)
			userRoutes.PUT("/update/:id", api_authed_user.UpdateUserInformation)

		}
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/register", api_public_auth.Register)
			authRoutes.POST("/login", api_public_auth.Authentication)
		}
	}

	return r
}
