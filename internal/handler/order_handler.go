package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"ayzhunis/hot-coffee/aerrors"
	"ayzhunis/hot-coffee/internal/service"
	"ayzhunis/hot-coffee/models"
	"ayzhunis/hot-coffee/utils"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		if errors.Is(err, aerrors.ErrExist) {
			utils.RespondWithError(w, http.StatusConflict, aerrors.ErrExist.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	slog.Info("Order created", utils.PostGroup())
	utils.RespondWithJSON(w, http.StatusCreated, order)
}

// get all orders
func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.GetOrders()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Orders received", utils.ReqGroup())
	utils.RespondWithJSON(w, http.StatusOK, orders)
}

// get order by id
func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		if errors.Is(err, aerrors.ErrNotExist) {
			utils.RespondWithError(w, http.StatusNotFound, aerrors.ErrNotExist.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	slog.Info("Order received by id", utils.ReqGroup())
	utils.RespondWithJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.orderService.UpdateOrder(&order, id); err != nil {
		if errors.Is(err, aerrors.ErrNotExist) {
			utils.RespondWithError(w, http.StatusNotFound, aerrors.ErrNotExist.Error())
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Order updated", utils.PutGroup())
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "updated successfully"})
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	if err := h.orderService.DeleteOrder(id); err != nil {
		if errors.Is(err, aerrors.ErrNotExist) {
			utils.RespondWithError(w, http.StatusNotFound, aerrors.ErrNotExist.Error())
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Order deleted", utils.DeleteGroup())
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Order deleted successfully"})
}

func (h *OrderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	if err := h.orderService.CloseOrder(id); err != nil {
		if errors.Is(err, aerrors.ErrNotExist) {
			utils.RespondWithError(w, http.StatusNotFound, aerrors.ErrNotExist.Error())
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Order closed", utils.PostGroup())
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Order closed successfully"})
}
