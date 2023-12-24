package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mealibek/gin-jwt/db"
)

func GetJwtUser(c *gin.Context) (db.User, error) {
	userInterface, exists := c.Get("user")
	if !exists {
		return db.User{}, errors.New("user not found in context")
	}

	// Type assertion to check if the user is of type db.User
	user, ok := userInterface.(db.User)
	if !ok {
		return db.User{}, errors.New("invalid user type in context")
	}

	return user, nil
}
