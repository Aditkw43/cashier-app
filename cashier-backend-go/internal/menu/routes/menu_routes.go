package routes

import (
	"cashier-backend-go/internal/menu/controller"
	"cashier-backend-go/middleware"

	"github.com/gin-gonic/gin"
)

type MenuRoutes struct {
	menuController *controller.MenuController
}

func NewMenuRoutes(menuController *controller.MenuController) *MenuRoutes {
	return &MenuRoutes{menuController: menuController}
}

func (r *MenuRoutes) MenuRoutes(router *gin.Engine) {
	menuGroup := router.Group("/menus")
	menuGroup.Use(middleware.JWTAuthMiddleware())
	{
		menuGroup.POST("/", r.menuController.CreateMenu)
		menuGroup.PUT("/:id", r.menuController.UpdateMenu)
		menuGroup.DELETE("/:id", r.menuController.DeleteMenu)
		menuGroup.GET("/:id", r.menuController.GetMenuByID)
		menuGroup.GET("/", r.menuController.GetAllMenus)
	}
}
