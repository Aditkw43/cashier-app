package repository

import (
	"cashier-backend-go/internal/domain/menu/model"

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
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) Create(menu *model.Menu) error {
	return r.db.Create(menu).Error
}

func (r *menuRepository) Update(menu *model.Menu) error {
	return r.db.Save(menu).Error
}

func (r *menuRepository) Delete(id uint) error {
	return r.db.Delete(&model.Menu{}, id).Error
}

func (r *menuRepository) FindByID(id uint) (*model.Menu, error) {
	var menu model.Menu
	if err := r.db.First(&menu, id).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) FindAll() ([]model.Menu, error) {
	var menus []model.Menu
	if err := r.db.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}
