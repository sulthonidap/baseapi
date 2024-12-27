package middleware

import (
	"baseApi/config"
	"baseApi/database"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginBody struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Active   bool   `json:"active"`
	ForceCpw bool   `json:"forceCpw"`
}

type AuthOptions struct {
	Roles []string
}

func Auth(options *AuthOptions) *jwt.GinJWTMiddleware {
	cfg := config.GetAll()

	var userId uint

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "baseApi API",
		Key:         []byte(cfg.Jwt.Secret),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					"id":       v.ID,
					"name":     v.Name,
					"email":    v.Email,
					"role":     v.Role,
					"forceCpw": v.ForceCpw,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id, _ := claims["id"].(float64)

			return &User{
				ID:       uint(id),
				Name:     claims["name"].(string),
				Email:    claims["email"].(string),
				Role:     claims["role"].(string),
				ForceCpw: claims["forceCpw"].(bool),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var user User
			var values loginBody

			if err := c.ShouldBind(&values); err != nil {
				log.Println(err)
				return "", jwt.ErrMissingLoginValues
			}

			result := database.DBConn.Where("username = ? AND active = TRUE", values.Username).First(&user)
			if result.Error != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			userId = user.ID

			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(values.Password))
			if err != nil {
				log.Println(err)
				return nil, jwt.ErrFailedAuthentication
			}

			var userData database.User
			database.DBConn.Where("id = ?", userId).Find(&userData)

			// Get Current Date
			current := time.Now()
			currentNano := current.UnixNano()

			// Get Last Year Update
			lastCpw := userData.LastCpw.AddDate(1, 0, 0)
			lastCpwNano := lastCpw.UnixNano()

			updateData := make(map[string]interface{})
			if currentNano > lastCpwNano {
				updateData["ForceCpw"] = true
			}
			database.DBConn.Model(&userData).Updates(updateData)
			return &user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			user, ok := data.(*User)
			if !ok {
				return false
			}

			if len(options.Roles) > 0 {
				for _, role := range options.Roles {
					if role == user.Role {
						return true
					}
				}

				return false
			}

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, t time.Time) {
			var user database.User
			database.DBConn.First(&user, userId)

			c.JSON(http.StatusOK, gin.H{
				"code":     code,
				"message":  "login successfully",
				"token":    token,
				"expire":   t.Format(time.RFC3339),
				"id":       user.ID,
				"username": user.Username,
				"fullname": user.Fullname,
				"role":     user.Role,
				"active":   user.Active,
				"forceCpw": user.ForceCpw,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Println(err)
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
