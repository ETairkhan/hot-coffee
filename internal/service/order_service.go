package service

import (
	"errors"
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
)

// OrderService defines the service that manages orders
type OrderService struct {
	repo dal.Repository
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(repo dal.Repository) *OrderService {
	return &OrderService{repo: repo}
}

// CreateOrder saves a new order
func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.repo.SaveOrder(order)
}

// GetOrders retrieves all orders
func (s *OrderService) GetOrders() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}

// GetOrderByID fetches a specific order by its ID
func (s *OrderService) GetOrderByID(id string) (*models.Order, error) {
	order, err := s.repo.GetOrderByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}

// UpdateOrder modifies an existing order
func (s *OrderService) UpdateOrder(order *models.Order) error {
	return s.repo.UpdateOrder(order)
}

// DeleteOrder removes an order by its ID
func (s *OrderService) DeleteOrder(id string) error {
	return s.repo.DeleteOrder(id)
}

// CloseOrder marks an order as "closed"
func (s *OrderService) CloseOrder(id string) error {
	order, err := s.repo.GetOrderByID(id)
	if err != nil {
		return errors.New("order not found")
	}

	order.Status = "closed"
	return s.repo.UpdateOrder(order)
}
