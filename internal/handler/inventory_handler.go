package handler

import (
	"ayzhunis/hot-coffee/aerrors"
	"ayzhunis/hot-coffee/internal/service"
	"ayzhunis/hot-coffee/models"
	"ayzhunis/hot-coffee/utils"
	"encoding/json"
	"errors"
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
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.inventoryService.CreateInventoryItems(&inv); err != nil {
		if errors.Is(err, aerrors.ErrExist) {
			utils.RespondWithError(w, http.StatusConflict, aerrors.ErrExist.Error())
			return
		}
		slog.Error(err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Inventory created", utils.PostGroup())
	utils.RespondWithJSON(w, http.StatusCreated, inv)
}

// get all inventory items
func (h *InventoryHandler) GetAllInventory(w http.ResponseWriter, r *http.Request) {
	invenItems, err := h.inventoryService.GetAllInventory()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Inventory items received", utils.ReqGroup())
	utils.RespondWithJSON(w, http.StatusOK, invenItems)
}

// get inventory item by id
func (h *InventoryHandler) GetInventoryById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing inventory ID")
		return
	}

	invenItems, err := h.inventoryService.GetInventoryById(id)
	if err != nil {
		if errors.Is(err, aerrors.ErrNotExist) {
			utils.RespondWithError(w, http.StatusNotFound, aerrors.ErrNotExist.Error())
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Inventory received by id", utils.ReqGroup())
	utils.RespondWithJSON(w, http.StatusOK, invenItems)
}

func (h *InventoryHandler) UpdateInventoryItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing inventory ID")
		return
	}

	var inv models.InventoryItem
	if err := json.NewDecoder(r.Body).Decode(&inv); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.inventoryService.UpdateInventoryItem(&inv, id); err != nil {
		if errors.Is(err, aerrors.ErrNotExist) {
			utils.RespondWithError(w, http.StatusNotFound, aerrors.ErrNotExist.Error())
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Inventory updated", utils.PutGroup())
	utils.RespondWithJSON(w, http.StatusOK, inv)
}

func (h *InventoryHandler) DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing inventory ID")
		return
	}

	if err := h.inventoryService.DeleteInventoryItem(id); err != nil {
		if errors.Is(err, aerrors.ErrNotExist) {
			utils.RespondWithError(w, http.StatusNotFound, aerrors.ErrNotExist.Error())
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Inventory deleted", utils.DeleteGroup())
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Inventory deleted successfully"})
}
