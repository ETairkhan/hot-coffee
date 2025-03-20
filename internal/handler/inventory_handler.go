package handler

import (
	"ayzhunis/hot-coffee/internal/service"
	"ayzhunis/hot-coffee/models"
	"ayzhunis/hot-coffee/utils"
	"encoding/json"
	"log/slog"
	"net/http"
)

type InvenotryHandler struct {
	inventoryService *service.InventoryService
}

func NewInventoryHandler(inventoryService *service.InventoryService) *InvenotryHandler {
	return &InvenotryHandler{inventoryService: inventoryService}
}

func (h *InvenotryHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	var inv models.InventoryItem
	if err := json.NewDecoder(r.Body).Decode(&inv); err != nil {
		slog.Error(err.Error())
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.inventoryService.CreateInve(&inv); err != nil {
		slog.Error(err.Error())
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Menu created", utils.PostGroup())
	h.respondWithJSON(w, http.StatusCreated, menu)
}

// get all menu
func (h *InvenotryHandler) GetAllMenuItems(w http.ResponseWriter, r *http.Request) {
	menuItems, err := h.menuService.GetAllMenu()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Menu items received", utils.ReqGroup())
	h.respondWithJSON(w, http.StatusOK, menuItems)
}

// get menu by id
func (h *InvenotryHandler) GetMenuItemByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	menuItems, err := h.menuService.GetMenuItemByID(id)
	if err != nil {
		h.respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	slog.Info("Menu received by id", utils.ReqGroup())
	h.respondWithJSON(w, http.StatusOK, menuItems)
}

func (h *InvenotryHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	var menu models.MenuItem
	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.menuService.UpdateMenuItem(&menu); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Menu updated", utils.PutGroup())
	h.respondWithJSON(w, http.StatusOK, menu)
}

func (h *InvenotryHandler) DeleteMenuItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	if err := h.menuService.DeleteMenuItemById(id); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Menu deleted", utils.DeleteGroup())
	h.respondWithJSON(w, http.StatusOK, map[string]string{"message": "Menu deleted successfully"})
}

func (h *InvenotryHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *InvenotryHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
