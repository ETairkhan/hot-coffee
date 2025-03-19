package service

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
	"errors"
)

// MenuService defines the service that manages orders
type MenuService struct {
	MenuRepo *dal.MenuItemsRepository
}

// NewMenuService creates a new instance of OrderService
func NewMenuService(repo *dal.MenuItemsRepository) *MenuService {
	return &MenuService{MenuRepo: repo}
}

// CreateOrder saves a new order
func (s *MenuService) CreateMenuItems(menu *models.MenuItem) error {
	return s.MenuRepo.CreateMenuItems(menu)
}

// GetOrders retrieves all orders
func (s *MenuService) GetAllMenu() (*[]models.MenuItem, error) {
	return s.MenuRepo.GetAllMenuItems()
}

// GetOrderByID fetches a specific order by its ID
func (s *MenuService) GetMenuItemByID(id string) (*models.MenuItem, error) {
	menu, err := s.MenuRepo.GetMenuItemByID(id)
	if err != nil {
		return nil, errors.New("menu not found")
	}
	return menu, nil
}

// UpdateOrder modifies an existing order
func (s *MenuService) UpdateMenuItem(menu *models.MenuItem) error {
	return s.MenuRepo.UpdateMenuItem(menu)
}

// DeleteOrder removes an order by its ID
func (s *MenuService) DeleteMenuItemById(id string) error {
	return s.MenuRepo.DeleteMenuItemById(id)
}
