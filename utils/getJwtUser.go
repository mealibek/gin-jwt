package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mealibek/gin-jwt/models"
)

func GetJwtUser(c *gin.Context) (models.User, error) {
	userInterface, exists := c.Get("user")
	if !exists {
		return models.User{}, errors.New("user not found in context")
	}

	// Type assertion to check if the user is of type models.User
	user, ok := userInterface.(models.User)
	if !ok {
		return models.User{}, errors.New("invalid user type in context")
	}

	return user, nil
}
