package user

import (
	"baseApi/database"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ViewUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Params.ByName("userId"))
	var user database.User
	err := database.DBConn.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
