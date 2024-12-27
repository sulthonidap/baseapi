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

// User form
type userForm struct {
	ID       string `form:"id" json:"id"`
	Username string `form:"username" json:"username" binding:"required"`
	Fullname string `form:"fullname" json:"fullname" binding:"required"`
	Active   bool   `form:"active" json:"active"`
	Password string `form:"password" json:"password"`
	Role     string `form:"role" json:"role" binding:"required"`
	ForceCpw string `form:"forceCpw" json:"forceCpw"`
}

// AddUser add a user
func AddUser(c *gin.Context) {
	var form userForm
	if err := c.ShouldBind(&form); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Only admin are allowed
	claims := jwt.ExtractClaims(c)
	if claims["role"] != "admin" {
		log.Println("Operation not allowed")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operation not allowed"})
		return
	}

	// Check operation
	var err error
	var ops string
	if form.ID != "" && form.ID != "0" {
		// Edit mode
		ops = "Update a user data: " + form.ID
		var user database.User
		database.DBConn.Where("id = ?", form.ID).Find(&user)
		updateData := make(map[string]interface{})

		if form.Username != "" {
			updateData["Username"] = form.Username
		}
		if form.Fullname != "" {
			updateData["Fullname"] = form.Fullname
		}
		if form.Role != "" {
			updateData["Role"] = form.Role
		}
		if form.ForceCpw == "active" {
			updateData["ForceCpw"] = true
		}
		if form.ForceCpw == "inactive" {
			updateData["ForceCpw"] = false
		}
		err = database.DBConn.Model(&user).Updates(&updateData).Error

	} else {
		// Insert mode
		ops = "Add a new user"
		// Validate username
		if isUserExist(form.Username) {
			log.Println("User already exist")
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "User already exist"})
			return
		}

		pass, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 14)
		current := time.Now()
		err = database.DBConn.Create(&database.User{
			Username: form.Username,
			Fullname: form.Fullname,
			Active:   true,
			Password: string(pass),
			Role:     form.Role,
			ForceCpw: form.ForceCpw == "active" && true,
			LastCpw:  &current,
		}).Error
	}

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Activity log
	database.AddActivitylog(
		claims["name"].(string),
		ops,
		c,
	)

	// Status toast
	c.JSON(http.StatusOK, gin.H{
		"message": "Add user success"})

}

func isUserExist(email string) bool {
	var count int64
	database.DBConn.Model(&database.User{}).Where(
		"email = ?", email).Count(&count)

	return count > 0
}
