package user

import (
	"baseApi/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListUser list all user
func ListUser(c *gin.Context) {
	var users []database.User
	role := c.Query("role")
	active := c.Query("active")

	sql := database.DBConn.Model(&database.User{})

	if role != "" {
		sql = sql.Where("role = ?", role)
	}

	if active == "active" {
		sql = sql.Where("active = TRUE")
	} else if active == "inactive" {
		sql = sql.Where("active = FALSE")
	}

	sql.Order("id desc").
		Find(&users)
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": users})
}
