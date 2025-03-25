package helper

import (
	"fmt"

	"ayzhunis/hot-coffee/aerrors"
	"ayzhunis/hot-coffee/models"
)

func CheckItemId(files []models.InventoryItem, id string) error {
	if files == nil {
		return fmt.Errorf("file is nil")
	}

	for _, items := range files {
		if items.IngredientID == id {
			return nil
		}
	}
	return fmt.Errorf("item with ID %s not found: %w", id, aerrors.ErrNotExist)
}

func CheckerForInventItems(inventItems *models.InventoryItem) error {
	if inventItems.IngredientID == "" {
		return fmt.Errorf("item ID should not be set when adding a new item")
	}
	if inventItems.Name == "" {
		return fmt.Errorf("please provide a name for the item")
	}

	if inventItems.Quantity <= 0 {
		return fmt.Errorf("please specify a quantity for the item")
	}

	if inventItems.Unit == "" {
		return fmt.Errorf("please provide a unit for the item")
	}

	return nil
}
