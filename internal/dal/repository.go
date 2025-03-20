package dal

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path"

	"ayzhunis/hot-coffee/models"
)

var (
	ErrClosedAlready  = errors.New("closed already")
	ErrStatusClosed   = errors.New("status is closed to change")
	ErrNotFound       = errors.New("not found")
	ErrDuplicateFound = errors.New("duplicate id found")
)

func GetAllItems[T models.Entity](dir, filename string) (*[]T, error) {
	var items []T
	data, err := os.ReadFile(path.Join(dir, filename))
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	return &items, nil
}

func GetById[T models.Entity](dir, filename, id string) (T, error) {
	var res T
	items, err := GetAllItems[T](dir, filename)
	if err != nil {
		return res, err
	}
	found := false
	for _, item := range *items {
		if item.GetID() == id {
			if found {
				return res, ErrDuplicateFound
			}
			found = true
			res = item
		}
	}

	if !found {
		return res, ErrNotFound
	}
	return res, nil
}

func CreateItem[T models.Entity](dir, filename string, item T) error {
	items, err := GetAllItems[T](dir, filename)
	if err != nil {
		return err
	}

	for i := range *items {
		if (*items)[i].GetID() == item.GetID() {
			return ErrDuplicateFound
		}
	}

	*items = append(*items, item)

	return writeItems(dir, filename, items)
}

func UpdateItem[T models.Entity](dir, filename string, item T) error {
	items, err := GetAllItems[T](dir, filename)
	if err != nil {
		return err
	}
	found := false
	for i := range *items {
		if (*items)[i].GetID() == item.GetID() {
			(*items)[i] = item
			if found {
				return ErrDuplicateFound
			}
			found = true
		}
	}

	if !found {
		return ErrNotFound
	}

	return writeItems(dir, filename, items)
} 

func DeleteItem[T models.Entity](dir, filename, id string) error {
	items, err := GetAllItems[T](dir, filename)
	if err != nil {
		return err
	}
	index := -1

	for i := range *items {
		if (*items)[i].GetID() == id {
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
	newItems := append((*items)[:index], (*items)[index+1:]...) // deleting element from array

	return writeItems(dir, filename, &newItems)
}

func writeItems[T models.Entity](dir, filename string, items *[]T) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(dir, filename), data, fs.FileMode(os.O_TRUNC))
	if err != nil {
		return err
	}
	return nil
}