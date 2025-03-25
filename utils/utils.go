package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ayzhunis/hot-coffee/models"
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
	slog.Warn("Responding with error", "code", code, "message", message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
