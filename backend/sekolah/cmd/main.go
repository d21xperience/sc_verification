// package main
package main

import (
	"net/http"
	"sekolah/controllers"
	"sekolah/routes"
	"time"

	"myproject/shared/db"

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

	// Inisialisasi database
	db.InitDatabase(dbConfig)
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

	var client = &http.Client{}
	dapoController := controllers.NewDapodikService(client)

	routes.SetupDapodik(router, dapoController)
	router.Run(":8081")

}
