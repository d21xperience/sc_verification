package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Mendapatkan aplikasi dapodik
func (ds DapodikClient) GetWSDapodik(ctx *gin.Context) {
	// sekolahID := ctx.Query("sekolah_id")
	fmt.Println(setDataDapo.BaseURL)
}
