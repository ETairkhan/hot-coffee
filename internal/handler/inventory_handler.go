package handler

import (
	"ayzhunis/hot-coffee/internal/service"
	"ayzhunis/hot-coffee/models"
	"ayzhunis/hot-coffee/utils"
	"encoding/json"
	"log/slog"
	"net/http"
)

type InventoryHandler struct {
	inventoryService *service.InventoryService
}

func NewInventoryHandler(inventoryService *service.InventoryService) *InventoryHandler {
	return &InventoryHandler{inventoryService: inventoryService}
}

func (h *InventoryHandler) CreateInventoryItems(w http.ResponseWriter, r *http.Request) {
	var inv models.InventoryItem
	if err := json.NewDecoder(r.Body).Decode(&inv); err != nil {
		slog.Error(err.Error())
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.inventoryService.CreateInventoryItems(&inv); err != nil {
		slog.Error(err.Error())
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Menu created", utils.PostGroup())
	h.respondWithJSON(w, http.StatusCreated, inv)
}

// get all menu
func (h *InventoryHandler) GetAllInventory(w http.ResponseWriter, r *http.Request) {
	invenItems, err := h.inventoryService.GetAllInventory()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Menu items received", utils.ReqGroup())
	h.respondWithJSON(w, http.StatusOK, invenItems)
}

// get menu by id
func (h *InventoryHandler) GetInventoryById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	invenItems, err := h.inventoryService.GetInventoryById(id)
	if err != nil {
		h.respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	slog.Info("Menu received by id", utils.ReqGroup())
	h.respondWithJSON(w, http.StatusOK, invenItems)
}

func (h *InventoryHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	var inv models.InventoryItem
	if err := json.NewDecoder(r.Body).Decode(&inv); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.inventoryService.UpdateInventoryItem(&inv); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Menu updated", utils.PutGroup())
	h.respondWithJSON(w, http.StatusOK, inv)
}

func (h *InventoryHandler) DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	if err := h.inventoryService.DeleteInventoryItem(id); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Menu deleted", utils.DeleteGroup())
	h.respondWithJSON(w, http.StatusOK, map[string]string{"message": "Menu deleted successfully"})
}

func (h *InventoryHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *InventoryHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
