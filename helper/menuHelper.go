package helper

import (
	"fmt"
	"log"

	"ayzhunis/hot-coffee/models"
	"ayzhunis/hot-coffee/aerrors"
)

func CheckerForMenuItems(items models.MenuItem, allInvnetory []models.InventoryItem) error {
	for _, inv := range allInvnetory {
		log.Printf("Inventory Item: %+v", inv)
	}

	if len(items.Ingredients) == 0 {
		return fmt.Errorf("menu item must have at least one ingredient")
	}

	for _, i := range items.Ingredients {
		if i.IngredientID == "" {
			return fmt.Errorf("ingredient ID should not be empty")
		}

		if err := CheckItemId(allInvnetory, i.IngredientID); err != nil {
			return err
		}

		if i.Quantity <= 0 {
			return fmt.Errorf("ingredient quantity should be greater than 0, got: ")
		}
	}
	return nil
}

func CheckMenuExistsId(files []models.MenuItem, id string) error {
	if files == nil {
		return fmt.Errorf("file is nils")
	}
	for _, i := range files {
		if i.ID == id {
			return nil
		}
	}
	return aerrors.ErrNotExist
}
