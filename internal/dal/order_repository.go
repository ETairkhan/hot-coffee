package dal

import (
	"ayzhunis/hot-coffee/models"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
)

var (
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
func (r *OrderRepository) GetAllOrders() (*[]models.Order, error) {
	orders := make([]models.Order, 0)
	fmt.Println(r.dir, orders)
	f, err := os.ReadFile(path.Join(r.dir, ordersFile))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(f, &orders); err != nil {
		return nil, err
	}

	return &orders, nil
}

// return only one order with id
func (r *OrderRepository) GetOrderByID(id string) (*models.Order, error) {
	orders := make([]models.Order, 0)
	res := models.Order{}

	f, err := os.ReadFile(path.Join(r.dir, ordersFile))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(f, &orders); err != nil {
		return nil, err
	}

	for _, order := range orders {
		if order.ID == id {
			res = order
		}
	}
	return &res, nil
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	orders := make([]models.Order, 0)

	f, err := os.ReadFile(path.Join(r.dir, ordersFile))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(f, &orders); err != nil {
		return err
	}
	orders = append(orders, *order)

	data, err := json.MarshalIndent(&orders, "", "  ") // create array of byte and contain with spaces
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(r.dir, ordersFile), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return nil
	}
	return nil
}
