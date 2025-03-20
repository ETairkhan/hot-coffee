package dal

import (
	"encoding/json"
	"io/fs"
	"os"
	"path"

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

func (r *OrderRepository) UpdateOrder(order *models.Order) error {
	orders := make([]models.Order, 0)

	f, err := os.ReadFile(path.Join(r.dir, ordersFile))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(f, &orders); err != nil {
		return err
	}
	found := false
	for i, ord := range orders {
		if order.ID == ord.ID {
			if order.Status == closed {
				return ErrStatusClosed
			}
			orders[i] = *order
			found = true
		}
	}
	if !found {
		return ErrNotFound
	}
	data, err := json.MarshalIndent(&orders, "", "  ") // create array of byte and contain spaces
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(r.dir, ordersFile), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) DeleteOrderById(id string) error {
	index := -1
	orders := make([]models.Order, 0)

	f, err := os.ReadFile(path.Join(r.dir, ordersFile))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(f, &orders); err != nil {
		return err
	}

	for i := range orders {
		if orders[i].ID == id {
			if index == -1 {
				index = i
			} else {
				return ErrDuplicateFound
			}
		}
	}
	if index < 0 {
		return ErrNotFound
	}
	newOrders := append(orders[:index], orders[index+1:]...) // deleting element from array
	data, err := json.MarshalIndent(&newOrders, "", "  ")    // create array of byte and contain spaces
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(r.dir, ordersFile), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) CloseOrder(id string) error {
	orders := make([]models.Order, 0)

	f, err := os.ReadFile(path.Join(r.dir, ordersFile))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(f, &orders); err != nil {
		return err
	}
	found := false

	for i := range orders {
		if orders[i].ID == id {
			if orders[i].Status == closed {
				return ErrClosedAlready
			}
			orders[i].Status = closed
			if found {
				return ErrDuplicateFound
			}
			found = true
		}
	}
	if !found {
		return ErrNotFound
	}
	data, err := json.MarshalIndent(&orders, "", "  ") // create array of byte and contain spaces
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(r.dir, ordersFile), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return err
	}
	return nil
}
