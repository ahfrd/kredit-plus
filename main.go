package main

import (
	"kredit-plus/konsumen/controller"
	"kredit-plus/konsumen/repository"
	"kredit-plus/konsumen/service"
	"kredit-plus/routes"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	env := helpers.Env{}
// 	env.StartingCheck()

// }
func main() {
	router := gin.Default()
	konsumenRepository := repository.NewKonsumenRepository()
	//Service
	konsumenService := service.NewKonsumenService(&konsumenRepository)
	//Controller
	konsumenController := controller.NewKonsumenController(&konsumenService)
	routes.SetupRoutes(router, &konsumenController)
	router.Run(":8000")
}
