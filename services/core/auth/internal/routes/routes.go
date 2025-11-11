package routes

import (
	"github.com/gin-gonic/gin"
	"auth/internal/controller"
	"auth/internal/service"
)

func RegisterRoutes(r *gin.Engine) {
	authService := service.AuthService{}
	authController := controller.AuthController{Service: authService}

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
}
