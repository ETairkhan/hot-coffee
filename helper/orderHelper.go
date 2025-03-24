package helper

import (
	"fmt"
	"strings"
	"ayzhunis/hot-coffee/aerrors"
	"ayzhunis/hot-coffee/models"
)

func CheckForOrders(orderItems models.Order, files []models.MenuItem) error {
	if orderItems.ID != "" {
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

	return nil
}

func CheckCustomerExists(files []models.Order, customerName string) error {
	if files == nil {
		return fmt.Errorf("file is nils")
	}

	for _, i := range files {
		if strings.ToLower(i.CustomerName) == strings.ToLower(customerName) {
			return nil
		}
	}
	return aerrors.ErrNotExist
}

func CheckCustomerExistsId(files []models.Order, id string) error {
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

func CheckStatus(files []models.Order, id string) error {
	if files == nil {
		return fmt.Errorf("file is nils")
	}
	for _, i := range files {
		if i.ID == id {
			if i.Status == "closed" {
				return fmt.Errorf("order is closed")
			}
		}
	}
	return nil
}
