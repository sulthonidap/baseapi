package auth

import (
	"baseApi/database"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type profileItemWeb struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Active   bool   `json:"active"`
	Password string `json:"password"`
	Role     string `json:"role"`
	ForceCpw string `json:"forceCpw"`
}

func Profile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	uid := claims["id"]
	var data profileItemWeb
	sqlQuery := `SELECT * FROM users WHERE id = ?`
	database.DBConn.Raw(sqlQuery, uid).Scan(&data)
	log.Println(data)

	c.JSON(200, gin.H{
		"data": data,
	})
}
