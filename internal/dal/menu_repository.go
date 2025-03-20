package dal

import (
	"encoding/json"
	"io/fs"
	"os"
	"path"

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
	index := -1
	menuItems := make([]models.MenuItem, 0)

	f, err := os.ReadFile(path.Join(r.dir, menuItemsFile))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(f, &menuItems); err != nil {
		return err
	}

	for i := range menuItems {
		if menuItems[i].ID == id {
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
	newMenuItems := append(menuItems[:index], menuItems[index+1:]...) // deleting element from array
	data, err := json.MarshalIndent(&newMenuItems, "", "  ")          // create array of byte and contain spaces
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(r.dir, menuItemsFile), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return err
	}
	return nil
}
