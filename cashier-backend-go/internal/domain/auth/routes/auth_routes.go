package routes

import (
	"cashier-backend-go/internal/domain/auth/controller"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authController *controller.AuthController
}

func NewAuthRoutes(authController *controller.AuthController) *AuthRoutes {
	return &AuthRoutes{authController: authController}
}

func (r *AuthRoutes) AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", r.authController.Login)
	}
}
