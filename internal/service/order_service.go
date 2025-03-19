package service

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
	"errors"
)

type OrderService struct {
	OrderRepo *dal.OrderRepository
}


func NewOrderService(repo *dal.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.OrderRepo.CreateOrder(order)
}

func (s *OrderService) GetOrders() (*[]models.Order, error) {
	return s.OrderRepo.GetAllOrders()
}

func (s *OrderService) GetOrderByID(id string) (*models.Order, error) {
	order, err := s.OrderRepo.GetOrderByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (s *OrderService) UpdateOrder(order *models.Order) error {
	return s.OrderRepo.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(id string) error {
	return s.OrderRepo.DeleteOrderById(id)
}

func (s *OrderService) CloseOrder(id string) error {
	return s.OrderRepo.CloseOrder(id)
}
