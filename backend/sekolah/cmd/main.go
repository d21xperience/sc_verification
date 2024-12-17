// package main
package main

import (
	"net/http"
	"sekolah/controllers"
	"sekolah/routes"
	"time"

	"myproject/shared/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Konfigurasi database dengan connection pooling
	dbConfig := db.Config{
		Host:            "localhost",
		Port:            5432,
		User:            "postgres",
		Password:        "",
		DBName:          "seckolahdb",
		MaxIdleConns:    10,               // Maksimum 10 koneksi idle
		MaxOpenConns:    100,              // Maksimum 100 koneksi total
		ConnMaxLifetime: 30 * time.Minute, // Koneksi direset setiap 30 menit
	}

	// var dataLogin service.DataDapo
	// dataLogin.BaseURL = "http://localhost:5774"
	// dataLogin.Username = "d21xperience@gmail.com"
	// dataLogin.Password = "Pasja123*"
	// dataLogin.SemeterID = "20222"
	// dataLogin.SekolahID = "8a5bd957-66bc-4096-9ff1-fee096b87ba0"

	// var client = &http.Client{}
	// dapo := service.NewdapodikServiceImpl(client)
	// dapo.LoginToDapodik(dataLogin)

	// // GET Bearer Token
	// tes := dapo.GetBearer(dataLogin.SekolahID)
	// fmt.Println(tes)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// router.GET("/api/v1/dapodik", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "CORS berhasil!"})
	// })
	// router.OPTIONS("/*path", func(c *gin.Context) {
	// 	c.Status(204) // Menjawab dengan status 204 No Content
	// })
	// Inisialisasi database
	db.InitDatabase(dbConfig)
	var client = &http.Client{}
	dapoController := controllers.NewDapodikClient(client)

	routes.SetupDapodik(router, dapoController)
	router.Run(":8082")

}
