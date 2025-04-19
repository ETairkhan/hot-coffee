package dal

import (
	"ayzhunis/hot-coffee/models"
)

const (
	closed     = "closed"
	ordersFile = "orders.json"
)

type OrderRepository struct {
	dir string
}

func NewOrderRepository(dir string) *OrderRepository {
	return &OrderRepository{
		dir: dir,
	}
}

// return all order which contains in data
func (r *OrderRepository) GetAllOrders() (*[]*models.Order, error) {
	return GetAllItems[*models.Order](r.dir, ordersFile)
}

// return only one order with id
func (r *OrderRepository) GetOrderByID(id string) (*models.Order, error) {
	return GetById[*models.Order](r.dir, ordersFile, id)
}

// create order by model
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return CreateItem(r.dir, ordersFile, order)
}

func (r *OrderRepository) UpdateOrder(order *models.Order, id string) error {
	return UpdateItem(r.dir, string(ordersFile), order, id)
}

func (r *OrderRepository) DeleteOrderById(id string) error {
	return DeleteItem[*models.Order](r.dir, ordersFile, id)
}

func (r *OrderRepository) CloseOrder(id string) error {
	orders, err := GetAllItems[*models.Order](r.dir, ordersFile)
	if err != nil {
		return err
	}
	found := false

	for i := range *orders {
		if (*orders)[i].ID == id {
			if (*orders)[i].Status == closed {
				return ErrClosedAlready
			}
			(*orders)[i].Status = closed
			if found {
				return ErrDuplicateFound
			}
			found = true
		}
	}
	if !found {
		return ErrNotFound
	}
	return writeItems(r.dir, string(ordersFile), orders)
}
