package service

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
)

type InventoryService struct {
	InvenRepo *dal.InventoryRepository
}

func NewInventoryService(repo *dal.InventoryRepository) *InventoryService {
	return &InventoryService{InvenRepo: repo}
}

func (s *InventoryService) CreateInventoryItems(menu *models.InventoryItem) error {
	return s.InvenRepo.CreateInventoryItems(menu)
}

func (s *InventoryService) GetAllMenu() (*[]models.MenuItem, error) {
	items, err := s.InvenRepo.GetAllMenuItems()
	if err != nil {
		return nil, err
	}

	newItems := make([]models.MenuItem, len(*items))
	for i := range *items {
		newItems[i] = *(*items)[i]
	}
	return &newItems, nil
}

func (s *InventoryService) GetMenuItemByID(id string) (*models.MenuItem, error) {
	menu, err := s.MenuRepo.GetMenuItemByID(id)
	if err != nil {
		return nil, errors.New("menu not found")
	}
	return menu, nil
}

func (s *InventoryService) UpdateMenuItem(menu *models.MenuItem) error {
	return s.MenuRepo.UpdateMenuItem(menu)
}

func (s *InventoryService) DeleteMenuItemById(id string) error {
	return s.MenuRepo.DeleteMenuItemById(id)
}
