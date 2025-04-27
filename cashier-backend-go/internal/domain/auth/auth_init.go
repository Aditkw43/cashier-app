package auth

import (
	"cashier-backend-go/internal/domain/auth/controller"
	"cashier-backend-go/internal/domain/auth/repository"
	"cashier-backend-go/internal/domain/auth/routes"
	"cashier-backend-go/internal/domain/auth/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuth(DB *gorm.DB, r *gin.Engine) {
	authRepo := repository.NewAuthRepository(DB)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	routes.NewAuthRoutes(authController).AuthRoutes(r)
}
