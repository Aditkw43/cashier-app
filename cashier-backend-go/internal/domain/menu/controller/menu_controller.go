package controller

import (
	"cashier-backend-go/internal/common/request"
	"cashier-backend-go/internal/common/response"
	"cashier-backend-go/internal/domain/menu/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	menuService service.MenuService
}

func NewMenuController(menuService service.MenuService) *MenuController {
	return &MenuController{menuService: menuService}
}

func (c *MenuController) CreateMenu(ctx *gin.Context) {
	var input request.CreateMenuRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid input",
			Error:   err.Error(),
		})
		return
	}

	menu, err := c.menuService.CreateMenu(input.Name, input.Price, input.Description, input.Category, input.Active)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create menu",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.ApiResponse{
		Status:  http.StatusCreated,
		Message: "Menu created successfully",
		Data:    menu,
	})
}

func (c *MenuController) UpdateMenu(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid menu ID",
		})
		return
	}

	var input request.UpdateMenuRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid input",
			Error:   err.Error(),
		})
		return
	}

	menu, err := c.menuService.UpdateMenu(uint(id), input.Name, input.Price, input.Description, input.Category, input.Active)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update menu",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Status:  http.StatusOK,
		Message: "Menu updated successfully",
		Data:    menu,
	})
}

func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid menu ID",
		})
		return
	}

	err = c.menuService.DeleteMenu(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete menu",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Status:  http.StatusOK,
		Message: "Menu deleted successfully",
	})
}

func (c *MenuController) GetMenuByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid menu ID",
		})
		return
	}

	menu, err := c.menuService.GetMenuByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse{
			Status:  http.StatusNotFound,
			Message: "Menu not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Status:  http.StatusOK,
		Message: "Menu found",
		Data:    menu,
	})
}

func (c *MenuController) GetAllMenus(ctx *gin.Context) {
	menus, err := c.menuService.GetAllMenus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get menus",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Status:  http.StatusOK,
		Message: "Menus retrieved successfully",
		Data:    menus,
	})
}
