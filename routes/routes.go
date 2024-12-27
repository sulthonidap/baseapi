package routes

import (
	"net/http"

	"baseApi/handler/auth"
	"baseApi/handler/storage"
	"baseApi/handler/user"
	"baseApi/handler/welcome"
	"baseApi/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter prepare for gin engine
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Page not found"})
	})

	// Logged-in users
	authMiddleware := middleware.Auth(&middleware.AuthOptions{})

	// Superadmin users
	adminMiddleware := middleware.Auth(&middleware.AuthOptions{
		Roles: []string{"admin"},
	})

	router.GET("/", welcome.Welcome)
	router.POST("/auth/login", authMiddleware.LoginHandler)

	// Handler route to get file from storage
	router.GET("/get/:year/:month/:date/:objectKey", storage.GetLocalStorage)

	// Logged-in users
	router.Use(authMiddleware.MiddlewareFunc())
	{
		router.GET("/user/profile", user.ViewProfile)
		router.GET("/profile", auth.Profile)
		router.POST("/profile/cpw", user.ChangeUserPassword)
	}

	// Admin
	router.Use(adminMiddleware.MiddlewareFunc())
	{
		router.GET("/users", user.ListUser)
		router.POST("/user", user.AddUser) // Add and Update
		router.GET("/user/:userId", user.ViewUser)
		router.POST("/user/rpw", user.ResetUserPassword)
	}

	return router
}
