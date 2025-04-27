package controller

import (
	"cashier-backend-go/internal/common/request"
	"cashier-backend-go/internal/common/response"
	"cashier-backend-go/internal/domain/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) Login(ctx *gin.Context) {
	var loginReq request.LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Error:   err.Error(),
		})
		return
	}

	token, err := ctrl.authService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ApiResponse{
			Status:  http.StatusUnauthorized,
			Message: "Invalid username or password",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Status:  http.StatusOK,
		Message: "Login successful",
		Data:    map[string]string{"token": token},
	})
}
