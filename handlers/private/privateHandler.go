package PrivateHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mealibek/gin-jwt/utils"
)

func Validate(c *gin.Context) {
	user, err := utils.GetJwtUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
