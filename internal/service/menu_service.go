package service

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
	"errors"
)

type MenuService struct {
	MenuRepo *dal.MenuItemsRepository
}

func NewMenuService(repo *dal.MenuItemsRepository) *MenuService {
	return &MenuService{MenuRepo: repo}
}

func (s *MenuService) CreateMenuItems(menu *models.MenuItem) error {
	return s.MenuRepo.CreateMenuItems(menu)
}

func (s *MenuService) GetAllMenu() (*[]models.MenuItem, error) {
	items, err := s.MenuRepo.GetAllMenuItems()
	if err != nil {
		return nil, err
	}

	newItems := make([]models.MenuItem, len(*items))
	for i := range *items {
		newItems[i] = *(*items)[i]
	}
	return &newItems, nil
}

func (s *MenuService) GetMenuItemByID(id string) (*models.MenuItem, error) {
	menu, err := s.MenuRepo.GetMenuItemByID(id)
	if err != nil {
		return nil, errors.New("menu not found")
	}
	return menu, nil
}

func (s *MenuService) UpdateMenuItem(menu *models.MenuItem) error {
	return s.MenuRepo.UpdateMenuItem(menu)
}

func (s *MenuService) DeleteMenuItemById(id string) error {
	return s.MenuRepo.DeleteMenuItemById(id)
}
