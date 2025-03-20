package dal

import (
	"encoding/json"
	"io/fs"
	"os"
	"path"

	"ayzhunis/hot-coffee/models"
)

const inventoryItemsFile = "inventory.json"

type InventoryRepository struct {
	dir string
}

func NewInventoryRepository(dir string) *InventoryRepository {
	return &InventoryRepository{
		dir: dir,
	}
}

func (ir *InventoryRepository) CreateInventoryItems(inventoryItem *models.InventoryItem) error {
	inventoryItems := make([]models.InventoryItem, 0)
	f, err := os.ReadFile(path.Join(ir.dir, inventoryItemsFile))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(f, &inventoryItems); err != nil {
		return err
	}

	for _, item1 := range inventoryItems {
		if item1.IngredientID == inventoryItem.IngredientID {
			return ErrDuplicateFound
		}
	}

	inventoryItems = append(inventoryItems, *inventoryItem)

	data, err := json.MarshalIndent(&inventoryItems, "", "  ") // create array of byte and contain with spaces
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(ir.dir, inventoryItemsFile), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return err
	}
	return nil
}

func (ir *InventoryRepository) GetAllInventory() (*[]*models.InventoryItem, error) {
	return GetAllItems[*models.InventoryItem](ir.dir, inventoryItemsFile)
}

func (ir *InventoryRepository) GetInventoryById(id string) (*models.InventoryItem, error) {
	inventoryItems := make([]models.InventoryItem, 0)
	var res *models.InventoryItem = nil

	f, err := os.ReadFile(path.Join(ir.dir, menuItemsFile))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(f, &inventoryItems); err != nil {
		return nil, err
	}

	for _, inventoryItem := range inventoryItems {
		if inventoryItem.IngredientID == id {
			res = &inventoryItem
		}
	}
	if res == nil {
		return nil, ErrNotFound
	}
	return res, nil
}

func (ir *InventoryRepository) UpdateInventoryItem(item *models.InventoryItem) error {
	inventoryItems := make([]models.InventoryItem, 0)

	f, err := os.ReadFile(path.Join(ir.dir, inventoryItemsFile))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(f, &inventoryItems); err != nil {
		return err
	}
	found := false // to check if the id exist
	for i, ord := range inventoryItems {
		if item.IngredientID == ord.IngredientID {
			inventoryItems[i] = *item
			found = true
		}
	}
	if !found {
		return ErrNotFound
	}
	data, err := json.MarshalIndent(&inventoryItems, "", "  ") // create array of byte and contain spaces
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(ir.dir, inventoryItemsFile), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return err
	}
	return nil
}

func (ir *InventoryRepository) DeleteInventoryItem(id string) error {
	index := -1
	invetoryItems := make([]models.InventoryItem, 0)

	f, err := os.ReadFile(path.Join(ir.dir, inventoryItemsFile))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(f, &invetoryItems); err != nil {
		return err
	}

	for i := range invetoryItems {
		if invetoryItems[i].IngredientID == id {
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
	newMenuItems := append(invetoryItems[:index], invetoryItems[index+1:]...) // deleting element from array
	data, err := json.MarshalIndent(&newMenuItems, "", "  ")                  // create array of byte and contain spaces
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(ir.dir, inventoryItemsFile), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return err
	}
	return nil
}
