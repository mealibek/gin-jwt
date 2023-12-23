package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mealibek/gin-jwt/handlers"
	PrivateHandlers "github.com/mealibek/gin-jwt/handlers/private"
	"github.com/mealibek/gin-jwt/initializers"
	"github.com/mealibek/gin-jwt/middlewares"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	// Grouping all routes and organizing...
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signup", handlers.SignUp)
			auth.POST("/signin", handlers.SignIn)
		}

		private := api.Group("/private")
		{
			private.GET("/validate", middlewares.RequireAuth(), PrivateHandlers.Validate)
		}

	}

	r.Run(fmt.Sprintf("localhost:%v", os.Getenv("PORT")))
}
