package service

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
	"errors"
)

// OrderService defines the service that manages orders
type OrderService struct {
	OrderRepo *dal.OrderRepository
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(repo *dal.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: repo}
}

// CreateOrder saves a new order
func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.OrderRepo.CreateOrder(order)
}

// GetOrders retrieves all orders
func (s *OrderService) GetOrders() (*[]models.Order, error) {
	return s.OrderRepo.GetAllOrders()
}

// GetOrderByID fetches a specific order by its ID
func (s *OrderService) GetOrderByID(id string) (*models.Order, error) {
	order, err := s.OrderRepo.GetOrderByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}

// UpdateOrder modifies an existing order
func (s *OrderService) UpdateOrder(order *models.Order) error {
	return s.OrderRepo.UpdateOrder(order)
}

// DeleteOrder removes an order by its ID
func (s *OrderService) DeleteOrder(id string) error {
	return s.OrderRepo.DeleteOrderById(id)
}

// CloseOrder marks an order as "closed"
func (s *OrderService) CloseOrder(id string) error {
	return s.OrderRepo.CloseOrder(id)
}
