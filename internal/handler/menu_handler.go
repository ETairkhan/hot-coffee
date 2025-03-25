package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sort"

	"ayzhunis/hot-coffee/internal/service"
	"ayzhunis/hot-coffee/models"
	"ayzhunis/hot-coffee/utils"
)

type MenuHandler struct {
	menuService *service.MenuService
}

func NewMenuHandler(menuService *service.MenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

func (h *MenuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	var menu models.MenuItem
	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
		slog.Error(err.Error())
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.menuService.CreateMenuItems(&menu); err != nil {
		slog.Error(err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Menu created", utils.PostGroup())
	utils.RespondWithJSON(w, http.StatusCreated, menu)
}

// get all menu
func (h *MenuHandler) GetAllMenuItems(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sort")
	menuItems, err := h.menuService.GetAllMenu()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	switch sortBy {
	case "ASC", "asc":
		sort.Slice(*menuItems, func(i, j int) bool {
			return (*menuItems)[i].Price < (*menuItems)[j].Price
		})
	case "DESC", "desc":
		sort.Slice(*menuItems, func(i, j int) bool {
			return (*menuItems)[i].Price > (*menuItems)[j].Price
		})
	}
	slog.Info("Menu items received", utils.ReqGroup())
	utils.RespondWithJSON(w, http.StatusOK, menuItems)
}

// get menu by id
func (h *MenuHandler) GetMenuItemByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	menuItems, err := h.menuService.GetMenuItemByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	slog.Info("Menu received by id", utils.ReqGroup())
	utils.RespondWithJSON(w, http.StatusOK, menuItems)
}

func (h *MenuHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}
	var menu models.MenuItem
	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.menuService.UpdateMenuItem(&menu, id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Menu updated", utils.PutGroup())
	utils.RespondWithJSON(w, http.StatusOK, menu)
}

func (h *MenuHandler) DeleteMenuItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	if err := h.menuService.DeleteMenuItemById(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Menu deleted", utils.DeleteGroup())
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Menu deleted successfully"})
}
