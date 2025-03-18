package dal

import (
	"ayzhunis/hot-coffee/models"
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

var (
	invertory  = "inventory.json"
	menu_items = "menu_items.json"
	orders     = "orders.json"
)

type Repository struct {
	dir string
}

func NewRepository(dir string) *Repository {
	return &Repository{
		dir: dir,
	}
}

// return all order which contains in data
func (r *Repository) GetAllOrders() (*[]models.Order, error) {
	m := make([]models.Order, 0)

	f, err := ioutil.ReadFile(path.Join(r.dir, orders))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(f, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

// return only one order with id
func (r *Repository) GetOrderByID(id string) (*models.Order, error) {
	m := make([]models.Order, 0)
	res := models.Order{}

	f, err := ioutil.ReadFile(path.Join(r.dir, orders))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(f, &m); err != nil {
		return nil, err
	}

	for _, order := range m {
		if order.ID == id {
			res = order
		}
	}
	return &res, nil
}

func (r *Repository) CreateOrder(order *models.Order) error {
	m := make([]models.Order, 0)

	f, err := ioutil.ReadFile(path.Join(r.dir, orders))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(f, &m); err != nil {
		return err
	}
	m = append(m, *order)

	data, err := json.Marshal(&m)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(r.dir, orders), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return nil
	}
	return nil
}
