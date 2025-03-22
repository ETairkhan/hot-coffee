package service

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
	"errors"
	"fmt"
)

type OrderService struct {
	OrderRepo        *dal.OrderRepository
	MenuService      *MenuService
	InventoryService *InventoryService
}

func NewOrderService(repo *dal.OrderRepository, menuService *MenuService, inventoryService *InventoryService) *OrderService {
	return &OrderService{
		OrderRepo:        repo,
		MenuService:      menuService,
		InventoryService: inventoryService,
	}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	// 1. Получаем все позиции меню
	menuItems, err := s.MenuService.GetAllMenu()
	if err != nil {
		return err
	}

	// 2. Получаем инвентарь
	inventory, err := s.InventoryService.GetAllInventory()
	if err != nil {
		return err
	}

	// 3. Создаем мапу ингредиентов в инвентаре
	invMap := make(map[string]*models.InventoryItem)
	for i := range *inventory {
		item := &(*inventory)[i]
		invMap[item.IngredientID] = item
	}

	// 4. Подсчитываем, сколько ингредиентов нужно
	needed := make(map[string]float64)
	for _, item := range order.Items {
		var found bool
		for _, menu := range *menuItems {
			if menu.ID == item.ProductID {
				found = true
				for _, ing := range menu.Ingredients {
					needed[ing.IngredientID] += ing.Quantity * float64(item.Quantity)
				}
			}
		}
		if !found {
			return errors.New("product not found in menu: " + item.ProductID)
		}
	}

	// 5. Проверяем, хватает ли ингредиентов
	for id, qty := range needed {
		if invItem, ok := invMap[id]; !ok {
			return errors.New("ingredient not found in inventory: " + id)
		} else if invItem.Quantity < qty {
			return errors.New("insufficient inventory for ingredient '" + invItem.Name + "'. Required: " + fmt.Sprintf("%.2f", qty) + " " + invItem.Unit + ", Available: " + fmt.Sprintf("%.2f", invItem.Quantity))
		}
	}

	// 6. Списываем ингредиенты
	for id, qty := range needed {
		invMap[id].Quantity -= qty
		err := s.InventoryService.UpdateInventoryItem(invMap[id], id)
		if err != nil {
			return err
		}
	}

	// 7. Создаем заказ
	return s.OrderRepo.CreateOrder(order)
}

func (s *OrderService) GetOrders() (*[]models.Order, error) {
	items, err := s.OrderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	newItems := make([]models.Order, len(*items))
	for i := range *items {
		newItems[i] = *(*items)[i]
	}
	return &newItems, nil
}

func (s *OrderService) GetOrderByID(id string) (*models.Order, error) {
	order, err := s.OrderRepo.GetOrderByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (s *OrderService) UpdateOrder(order *models.Order, id string) error {
	return s.OrderRepo.UpdateOrder(order, id)
}

func (s *OrderService) DeleteOrder(id string) error {
	return s.OrderRepo.DeleteOrderById(id)
}

func (s *OrderService) CloseOrder(id string) error {
	return s.OrderRepo.CloseOrder(id)
}
