package routes

import (
	"sekolah/controllers"

	"github.com/gin-gonic/gin"
)

// func DapodikRoutes(router *gin.Engine) {
// 	// Setup Dapodik

// 	dapodikRoutes := router.Group("/WebService")
// 	{
// 		dapodikRoutes.GET("/getSekolah", controllers.GetPesertaDidik())
// 		dapodikRoutes.GET("/getRombonganBelajar")
// 		dapodikRoutes.GET("/getPesertaDidik")
// 	}
// }

func SetupDapodik(router *gin.Engine, dapoController *controllers.DapodikService) {
	dapodikRouter := router.Group("/dapodik")
	{
		dapodikRouter.GET("/", dapoController.GetDapodikApp)
		dapodikRouter.POST("/login", dapoController.LoginToDapodik)
		dapodikRouter.GET("/GetPesertaDidik", dapoController.GetPesertaDidik)
	}
}
