package service

import (
	"cashier-backend-go/internal/domain/menu/model"
	"cashier-backend-go/internal/domain/menu/repository"
	"errors"
)

type MenuService interface {
	CreateMenu(name string, price float64, description string, category string, isActive bool) (*model.Menu, error)
	UpdateMenu(id uint, name string, price float64, description string, category string, isActive bool) (*model.Menu, error)
	DeleteMenu(id uint) error
	GetMenuByID(id uint) (*model.Menu, error)
	GetAllMenus() ([]model.Menu, error)
}

type menuService struct {
	menuRepo repository.MenuRepository
}

func NewMenuService(menuRepo repository.MenuRepository) MenuService {
	return &menuService{menuRepo: menuRepo}
}

func (s *menuService) CreateMenu(name string, price float64, description string, category string, isActive bool) (*model.Menu, error) {
	if name == "" || price <= 0 || category == "" {
		return nil, errors.New("invalid menu data")
	}
	menu := &model.Menu{Name: name, Price: price, Description: description, Category: category, IsActive: isActive}
	err := s.menuRepo.Create(menu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s *menuService) UpdateMenu(id uint, name string, price float64, description string, category string, isActive bool) (*model.Menu, error) {
	menu, err := s.menuRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	menu.Name = name
	menu.Price = price
	menu.Description = description
	menu.Category = category
	menu.IsActive = isActive

	err = s.menuRepo.Update(menu)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (s *menuService) DeleteMenu(id uint) error {
	_, err := s.menuRepo.FindByID(id)
	if err != nil {
		return err
	}

	return s.menuRepo.Delete(id)
}

func (s *menuService) GetMenuByID(id uint) (*model.Menu, error) {
	return s.menuRepo.FindByID(id)
}

func (s *menuService) GetAllMenus() ([]model.Menu, error) {
	return s.menuRepo.FindAll()
}
