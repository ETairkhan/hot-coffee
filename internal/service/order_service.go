package service

import (
	"errors"
	"fmt"

	"ayzhunis/hot-coffee/helper"
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
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

	// 1.5 Proverka ordera na validnation
	err = helper.CheckForOrders(*order, *menuItems)
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
	// 1. Get existing order
	existing, err := s.OrderRepo.GetOrderByID(id)
	if err != nil {
		return errors.New("order not found")
	}

	// 2. Get all menu items
	menuItems, err := s.MenuService.GetAllMenu()
	if err != nil {
		return err
	}

	// 3. Get inventory
	inventory, err := s.InventoryService.GetAllInventory()
	if err != nil {
		return err
	}

	// 4. Create inventory map
	invMap := make(map[string]*models.InventoryItem)
	for i := range *inventory {
		item := &(*inventory)[i]
		invMap[item.IngredientID] = item
	}

	// 5. Restore ingredients from the existing order
	for _, item := range existing.Items {
		for _, menu := range *menuItems {
			if menu.ID == item.ProductID {
				for _, ing := range menu.Ingredients {
					invMap[ing.IngredientID].Quantity += ing.Quantity * float64(item.Quantity)
				}
			}
		}
	}

	// 6. Calculate ingredients for new order
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

	// 7. Check inventory availability
	for id, qty := range needed {
		if invItem, ok := invMap[id]; !ok {
			return errors.New("ingredient not found in inventory: " + id)
		} else if invItem.Quantity < qty {
			return errors.New("insufficient inventory for ingredient '" + invItem.Name +
				"'. Required: " + fmt.Sprintf("%.2f", qty) + " " + invItem.Unit +
				", Available: " + fmt.Sprintf("%.2f", invItem.Quantity))
		}
	}

	// 8. Deduct new ingredients and update inventory
	for id, qty := range needed {
		invMap[id].Quantity -= qty
		err := s.InventoryService.UpdateInventoryItem(invMap[id], id)
		if err != nil {
			return err
		}
	}

	// 9. Update the order
	return s.OrderRepo.UpdateOrder(order, id)
}

func (s *OrderService) DeleteOrder(id string) error {
	// 1. Get the order
	order, err := s.OrderRepo.GetOrderByID(id)
	if err != nil {
		return errors.New("order not found")
	}

	// Only restore ingredients if order is still open
	if order.Status == "open" {
		// 2. Get all menu items
		menuItems, err := s.MenuService.GetAllMenu()
		if err != nil {
			return err
		}

		// 3. Get inventory
		inventory, err := s.InventoryService.GetAllInventory()
		if err != nil {
			return err
		}

		// 4. Create a map of inventory for fast access
		invMap := make(map[string]*models.InventoryItem)
		for i := range *inventory {
			item := &(*inventory)[i]
			invMap[item.IngredientID] = item
		}

		// 5. Loop through items in the order
		for _, item := range order.Items {
			for _, menu := range *menuItems {
				if menu.ID == item.ProductID {
					for _, ing := range menu.Ingredients {
						if invItem, exists := invMap[ing.IngredientID]; exists {
							// Add ingredients back based on quantity
							invItem.Quantity += ing.Quantity * float64(item.Quantity)

							// Update the inventory item
							err := s.InventoryService.UpdateInventoryItem(invItem, invItem.IngredientID)
							if err != nil {
								return fmt.Errorf("failed to restore inventory for ingredient %s: %v", ing.IngredientID, err)
							}
						}
					}
				}
			}
		}
	}

	// 6. Delete the order
	return s.OrderRepo.DeleteOrderById(id)
}

func (s *OrderService) CloseOrder(id string) error {
	return s.OrderRepo.CloseOrder(id)
}
