package handler

import (
	"ayzhunis/hot-coffee/internal/service"
	"ayzhunis/hot-coffee/models"
	"ayzhunis/hot-coffee/utils"
	"encoding/json"
	"log/slog"
	"net/http"
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
		slog.Error(err.Error())
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		slog.Error(err.Error())
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Order created", utils.PostGroup())
	h.respondWithJSON(w, http.StatusCreated, order)
}

// get all orders
func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.GetOrders()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Orders received", utils.ReqGroup())
	h.respondWithJSON(w, http.StatusOK, orders)
}

// get order by id
func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		h.respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	slog.Info("Order received by id", utils.ReqGroup())
	h.respondWithJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.orderService.UpdateOrder(&order); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Order updated", utils.PutGroup())
	h.respondWithJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	if err := h.orderService.DeleteOrder(id); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Order deleted", utils.DeleteGroup())
	h.respondWithJSON(w, http.StatusOK, map[string]string{"message": "Order deleted successfully"})
}

func (h *OrderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	if err := h.orderService.CloseOrder(id); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	slog.Info("Order closed", utils.PostGroup())
	h.respondWithJSON(w, http.StatusOK, map[string]string{"message": "Order closed successfully"})
}

func (h *OrderHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *OrderHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
