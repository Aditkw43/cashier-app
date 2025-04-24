package auth

import (
	"cashier-backend-go/internal/auth/controller"
	"cashier-backend-go/internal/auth/repository"
	"cashier-backend-go/internal/auth/routes"
	"cashier-backend-go/internal/auth/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuth(DB *gorm.DB, r *gin.Engine) {
	userRepo := repository.NewUserRepository(DB)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	routes.NewAuthRoutes(authController).AuthRoutes(r)
}
