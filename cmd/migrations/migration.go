package main

import (
	"github.com/mealibek/gin-jwt/db"
	"github.com/mealibek/gin-jwt/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&db.User{})
}
