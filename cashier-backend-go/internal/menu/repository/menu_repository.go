package repository

import (
	"cashier-backend-go/internal/menu/model"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu *model.Menu) error
	Update(menu *model.Menu) error
	Delete(id uint) error
	FindByID(id uint) (*model.Menu, error)
	FindAll() ([]model.Menu, error)
}

type menuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(DB *gorm.DB) MenuRepository {
	return &menuRepository{DB: DB}
}

func (r *menuRepository) Create(menu *model.Menu) error {
	return r.DB.Create(menu).Error
}

func (r *menuRepository) Update(menu *model.Menu) error {
	return r.DB.Save(menu).Error
}

func (r *menuRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Menu{}, id).Error
}

func (r *menuRepository) FindByID(id uint) (*model.Menu, error) {
	var menu model.Menu
	err := r.DB.First(&menu, id).Error
	return &menu, err
}

func (r *menuRepository) FindAll() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.DB.Find(&menus).Error
	return menus, err
}
