package user

import (
	"baseApi/database"
	"errors"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ViewProfile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	var user database.User
	err := database.DBConn.First(&user, claims["id"]).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		}
		return
	}

	var profileData struct {
		ID           uint       `json:"id"`
		Name         string     `json:"name"`
		Email        string     `json:"email"`
		Active       bool       `json:"active"`
		Role         string     `json:"role"`
		ForceCpw     bool       `json:"forceCpw"`
		LastCpw      *time.Time `json:"lastCpw"`
		UserEselonID *int       `json:"userEselonID"`
	}

	c.JSON(http.StatusOK, gin.H{
		"data": profileData,
	})
}
