package menu

import (
	"cashier-backend-go/internal/domain/menu/controller"
	"cashier-backend-go/internal/domain/menu/repository"
	"cashier-backend-go/internal/domain/menu/routes"
	"cashier-backend-go/internal/domain/menu/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitMenu(DB *gorm.DB, r *gin.Engine) {
	menuRepo := repository.NewMenuRepository(DB)
	menuService := service.NewMenuService(menuRepo)
	menuController := controller.NewMenuController(menuService)

	routes.NewMenuRoutes(menuController).MenuRoutes(r)
}
