package dal

import (
	"ayzhunis/hot-coffee/models"
)

const (
	menuItemsFile = "menu_items.json"
)

type MenuItemsRepository struct {
	dir string
}

func NewMenuRepository(dir string) *MenuItemsRepository {
	return &MenuItemsRepository{
		dir: dir,
	}
}

func (mr *MenuItemsRepository) GetAllMenuItems() (*[]*models.MenuItem, error) {
	return GetAllItems[*models.MenuItem](mr.dir, menuItemsFile)
}

func (r *MenuItemsRepository) GetMenuItemByID(id string) (*models.MenuItem, error) {
	return GetById[*models.MenuItem](r.dir, menuItemsFile, id)
}

// add menu to db if id duplicate error will return
func (mr *MenuItemsRepository) CreateMenuItems(item *models.MenuItem) error {
	return CreateItem(mr.dir, menuItemsFile, item)
}

func (r *MenuItemsRepository) UpdateMenuItem(item *models.MenuItem) error {
	return UpdateItem(r.dir, menuItemsFile, item)
}

// deleting or
func (r *MenuItemsRepository) DeleteMenuItemById(id string) error {
	return DeleteItem[*models.MenuItem](r.dir, menuItemsFile, id)
}
