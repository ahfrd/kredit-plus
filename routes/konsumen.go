package routes

import (
	"kredit-plus/konsumen/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, konsumenController *controller.KonsumenController) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/inquiry", konsumenController.Inquiry)
		v1.POST("/payment", konsumenController.Payment)
	}
}
