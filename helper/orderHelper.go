package helper

import (
	"fmt"

	"ayzhunis/hot-coffee/models"
)

func CheckForOrders(orderItems models.Order, files []models.MenuItem) error {
	if orderItems.ID == "" {
		return fmt.Errorf("order validation failed: order ID is already set: %s", orderItems.ID)
	}

	if orderItems.CustomerName == "" {
		return fmt.Errorf("please provide a customer name for the order")
	}

	for _, items := range orderItems.Items {
		if items.ProductID == "" {
			return fmt.Errorf("please provide a product ID for one of the items in the order")
		}
		if err := CheckMenuExistsId(files, items.ProductID); err != nil {
			return err
		}
		if items.Quantity <= 0 {
			return fmt.Errorf("please specify a quantity greater than zero for the item")
		}
	}
	if len(orderItems.Items) <= 0 {
		return fmt.Errorf("please provide items for the order")
	}
	if orderItems.Status == "open" || orderItems.Status == "closed" {
		return fmt.Errorf("please provide status")
	}
	return nil
}
