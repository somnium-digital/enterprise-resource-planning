package main

import (
	"github.com/gin-gonic/gin"
	"auth/config"
	"auth/internal/models"
	"auth/internal/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run() // port dari .env otomatis
}
