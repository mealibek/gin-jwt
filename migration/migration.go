package main

import (
	"github.com/mealibek/gin-jwt/initializers"
	"github.com/mealibek/gin-jwt/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
