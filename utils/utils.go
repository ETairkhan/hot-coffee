package utils

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"path"
	"regexp"

	"ayzhunis/hot-coffee/models"
)

const (
	inventory  = "inventory.json"
	menu_items = "menu_items.json"
	orders     = "orders.json"
)

func ReqGroup() slog.Attr {
	reqGroup := slog.Group(
		"request",
		"method", "GET",
	)
	return reqGroup
}

func PostGroup() slog.Attr {
	PostGroup := slog.Group(
		"request",
		"method", "POST",
	)
	return PostGroup
}

func PutGroup() slog.Attr {
	PostGroup := slog.Group(
		"request",
		"method", "PUT",
	)
	return PostGroup
}

func DeleteGroup() slog.Attr {
	PostGroup := slog.Group(
		"request",
		"method", "DELETE",
	)
	return PostGroup
}

func IsContain[T models.Entity](id string, items *[]T) (T, bool) {
	var res T
	for _, item := range *items {
		if (item).GetID() == id {
			return item, true
		}
	}
	return res, false
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	slog.Error(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func CheckDir(dirname string) error {
	re := regexp.MustCompile("^[a-zA-z0-9]+$")
	if !re.MatchString(dirname) {
		return errors.New("invalide filename")
	}

	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		if err := os.Mkdir(dirname, 0755); err != nil {
			return err
		}
	}
	files := []string{inventory, menu_items, orders}

	for _, f := range files {
		if err := CreateNotExist(dirname, f); err != nil {
			return err
		}
	}

	return nil
}

func CreateNotExist(dirname, filename string) error {
	if _, err := os.Stat(path.Join(dirname, filename)); os.IsNotExist(err) {
		// fmt.Println(path.Join(dirname, filename))
		f, err := os.Create(path.Join(dirname, filename))
		if err != nil {
			return err
		}
		_, err = f.Write([]byte(`[]`))
		if err != nil {
			return err
		}
	}
	return nil
}
