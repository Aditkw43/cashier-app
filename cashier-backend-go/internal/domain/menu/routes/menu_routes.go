package routes

import (
	"cashier-backend-go/internal/domain/menu/controller"

	"github.com/gin-gonic/gin"
)

type MenuRoutes struct {
	menuController *controller.MenuController
}

func NewMenuRoutes(menuController *controller.MenuController) *MenuRoutes {
	return &MenuRoutes{menuController: menuController}
}

func (r *MenuRoutes) MenuRoutes(router *gin.Engine) {
	menu := router.Group("/menus")
	{
		menu.POST("/", r.menuController.CreateMenu)
		menu.PUT("/:id", r.menuController.UpdateMenu)
		menu.DELETE("/:id", r.menuController.DeleteMenu)
		menu.GET("/:id", r.menuController.GetMenuByID)
		menu.GET("/", r.menuController.GetAllMenus)
	}
}
