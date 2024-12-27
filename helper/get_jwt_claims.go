package helper

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// GetJwtClaims will return user id and role, e.g. {"uid": 1, "role": "admin"}
func GetJwtClaims(c *gin.Context) interface{} {
	jwt.ExtractClaims(c)
	claims, _ := c.Get("id")

	return claims
}
