package service

import (
	"errors"
	"fmt"

	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/models"
	"ayzhunis/hot-coffee/utils"
)

const (
	open   = "open"
	closed = "closed"
)

var ErrNoItems = errors.New("there is no items")

type AggregationService struct {
	orderRepo     *dal.OrderRepository
	menuRepo      *dal.MenuItemsRepository
	inventoryRepo *dal.InventoryRepository
}

func NewAggregationService(
	orderRepo *dal.OrderRepository,
	menuRepo *dal.MenuItemsRepository,
	inventoryRepo *dal.InventoryRepository,
) *AggregationService {
	return &AggregationService{
		orderRepo:     orderRepo,
		menuRepo:      menuRepo,
		inventoryRepo: inventoryRepo,
	}
}

func (a *AggregationService) TotalSales() (float64, error) {
	var res float64 = 0

	orders, err := a.orderRepo.GetAllOrders()
	if err != nil {
		return 0, err
	}
	menus, err := a.menuRepo.GetAllMenuItems()
	if err != nil {
		return 0, err
	}

	for _, order := range *orders {
		if (*order).Status == closed {
			for _, item := range (*order).Items {
				obj, f := utils.IsContain(item.ProductID, menus)
				if !f {
					return 0, dal.ErrNotFound
				}
				res += obj.Price * float64(item.Quantity)
			}
		}
	}
	return res, nil
}

func (a *AggregationService) PopularItems() (*models.MenuItem, error) {
	var res *models.MenuItem
	id := ""
	mx := 0
	m := make(map[string]int, 0)

	orders, err := a.orderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	for _, order := range *orders {
		if (*order).Status == closed {
			for _, item := range (*order).Items {
				m[item.ProductID]++
				if m[item.ProductID] > mx {
					id = item.ProductID
					mx = m[item.ProductID]
				}
			}
		}
	}
	fmt.Println(m)
	if id == "" || mx == 0 {
		return nil, ErrNoItems
	}
	res, err = a.menuRepo.GetMenuItemByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
