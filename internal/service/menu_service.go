package service

import (
	"errors"

	"ayzhunis/hot-coffee/helper"
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
)

type MenuService struct {
	MenuRepo *dal.MenuItemsRepository
	IvenRepo *dal.InventoryRepository
}

func NewMenuService(repo *dal.MenuItemsRepository, iven *dal.InventoryRepository) *MenuService {
	return &MenuService{MenuRepo: repo}
}

func (s *MenuService) CreateMenuItems(menu *models.MenuItem) error {
	iventory, err := s.IvenRepo.GetAllInventory()
	if err != nil {
		return err
	}
	inve := make([]models.InventoryItem, len(*iventory))
	for i, s := range *iventory {
		inve[i] = *s
	}
	err = helper.CheckerForMenuItems(*menu, inve)
	if err != nil {
		return err
	}
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

func (s *MenuService) UpdateMenuItem(menu *models.MenuItem, id string) error {
	return s.MenuRepo.UpdateMenuItem(menu, id)
}

func (s *MenuService) DeleteMenuItemById(id string) error {
	return s.MenuRepo.DeleteMenuItemById(id)
}
