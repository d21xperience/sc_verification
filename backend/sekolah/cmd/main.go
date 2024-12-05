// package main
package main

import (
	"net/http"
	"sekolah/controllers"
	"sekolah/routes"

	"github.com/gin-gonic/gin"
)

func main() {

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
