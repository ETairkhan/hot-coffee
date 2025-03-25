package service

import (
	"errors"

	"ayzhunis/hot-coffee/helper"
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
)

type InventoryService struct {
	InvenRepo *dal.InventoryRepository
}

func NewInventoryService(repo *dal.InventoryRepository) *InventoryService {
	return &InventoryService{InvenRepo: repo}
}

func (s *InventoryService) CreateInventoryItems(iven *models.InventoryItem) error {
	err := helper.CheckerForInventItems(*iven)
	if err != nil {
		return err
	}

	return s.InvenRepo.CreateInventoryItem(iven)
}

func (s *InventoryService) GetAllInventory() (*[]models.InventoryItem, error) {
	items, err := s.InvenRepo.GetAllInventory()
	if err != nil {
		return nil, err
	}

	newItems := make([]models.InventoryItem, len(*items))
	for i := range *items {
		newItems[i] = *(*items)[i]
	}
	return &newItems, nil
}

func (s *InventoryService) GetInventoryById(id string) (*models.InventoryItem, error) {
	inven, err := s.InvenRepo.GetInventoryById(id)
	if err != nil {
		return nil, errors.New("inventory not found")
	}
	return inven, nil
}

func (s *InventoryService) UpdateInventoryItem(inven *models.InventoryItem, id string) error {
	err := helper.CheckerForInventItems(*inven)
	if err != nil {
		return err
	}
	return s.InvenRepo.UpdateInventoryItem(inven, id)
}

func (s *InventoryService) DeleteInventoryItem(id string) error {
	return s.InvenRepo.DeleteInventoryItem(id)
}
