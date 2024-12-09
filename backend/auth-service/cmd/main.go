package main

import (
	"auth-service/config"
	"auth-service/controller"
	"auth-service/repository"
	"auth-service/service"
	"auth-service/util"
	"time"

	"myproject/shared/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup database
	// Konfigurasi database dengan connection pooling
	dbConfig := db.Config{
		Host:            "localhost",
		Port:            5432,
		User:            "postgres",
		Password:        "",
		DBName:          "authdb",
		MaxIdleConns:    10,               // Maksimum 10 koneksi idle
		MaxOpenConns:    100,              // Maksimum 100 koneksi total
		ConnMaxLifetime: 30 * time.Minute, // Koneksi direset setiap 30 menit
	}

	// Inisialisasi database
	db.InitDatabase(dbConfig)

	// config.InitDatabase()
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
