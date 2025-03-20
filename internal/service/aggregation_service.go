package service

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/utils"
)

const (
	open = "open"
	closed = "closed"
)

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

func (a *AggregationService) TotalSales() (uint64, error) {
	var res uint64 = 0

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
				res += uint64(obj.Price)
			}
		}
	}
	return res, nil
}