package dal

import (
	"ayzhunis/hot-coffee/models"
	"encoding/json"
	"os"
	"path"
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
