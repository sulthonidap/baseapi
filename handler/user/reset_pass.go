package user

import (
	"baseApi/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type resetPasswordForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func ResetUserPassword(c *gin.Context) {
	// Prepare form
	var form resetPasswordForm
	if err := c.ShouldBind(&form); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Get the password
	var user database.User
	database.DBConn.
		Select("password, id").
		Where("id = ?", form.User).
		First(&user)

	if user.ID == 0 {
		log.Println("User not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Update the password
	pass, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 14)
	database.DBConn.Model(&user).Update("password", pass)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success reset password"})
}
