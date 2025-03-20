package dal

import (
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

func (ir *InventoryRepository) GetAllInventory() (*[]*models.InventoryItem, error) {
	return GetAllItems[*models.InventoryItem](ir.dir, inventoryItemsFile)
}

func (ir *InventoryRepository) GetInventoryById(id string) (*models.InventoryItem, error) {
	return GetById[*models.InventoryItem](ir.dir, inventoryItemsFile, id)
}

func (ir *InventoryRepository) CreateInventoryItem(item *models.InventoryItem) error {
	return CreateItem(ir.dir, inventoryItemsFile, item)
}

func (ir *InventoryRepository) UpdateInventoryItem(item *models.InventoryItem) error {
	return UpdateItem(ir.dir, inventoryItemsFile, item)
}

func (ir *InventoryRepository) DeleteInventoryItem(id string) error {
	return DeleteItem[*models.InventoryItem](ir.dir, inventoryItemsFile, id)
}
