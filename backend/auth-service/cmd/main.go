package main

import (
	"auth-service/config"
	"auth-service/controller"
	"auth-service/repository"
	"auth-service/service"
	"auth-service/util"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup database
	config.InitDatabase()
	// defer config.DB.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(config.DB)

	// Initialize services
	authService := service.NewAuthService(userRepo)

	// Initialize controllers
	authController := controller.NewAuthController(authService)

	// Setup Gin
	r := gin.Default()
	// Middleware
	authMiddleware := util.JWTAuthMiddleware()
	// Routes
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
	r.GET("/users/:id", authMiddleware, authController.GetUserByID)  // Ambil user berdasarkan ID
	r.GET("/profile", authMiddleware, authController.GetUserProfile) // Ambil profil user dari JWT
	r.POST("/upload/:id/profile_picture", authMiddleware, authController.UploadProfilePicture)

	// Start server
	r.Run(":8080")

}
