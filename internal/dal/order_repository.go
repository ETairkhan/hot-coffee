package dal

import (
	"ayzhunis/hot-coffee/models"
	"encoding/json"
	"io/ioutil"
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
