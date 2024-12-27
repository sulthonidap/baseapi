package user

import (
	"baseApi/database"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type changePasswordForm struct {
	OldPassword string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

func ChangeUserPassword(c *gin.Context) {
	// Prepare form
	var form changePasswordForm
	if err := c.ShouldBind(&form); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user
	claims := jwt.ExtractClaims(c)
	userId := claims["id"]

	// Get the password
	var user database.User
	database.DBConn.
		Select("password, id").
		Where("id = ?", userId).
		First(&user)

	if user.ID == 0 {
		log.Println("User not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.OldPassword))
	if err != nil {
		log.Println("Wrong password " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	current := time.Now()
	// Update the password
	pass, _ := bcrypt.GenerateFromPassword([]byte(form.NewPassword), 14)

	object := make(map[string]interface{})
	object["Password"] = string(pass)
	object["LastCpw"] = &current
	object["ForceCpw"] = false
	database.DBConn.Model(&user).Updates(&object)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success change password"})
}
