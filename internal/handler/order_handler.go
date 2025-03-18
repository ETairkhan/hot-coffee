package handler

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/internal/service"
	"ayzhunis/hot-coffee/models"
	"encoding/json"
	"log"
	"net/http"
)

type OrderHandler struct {
	orderService service.OrderService
	repo         *dal.Repository
}

func NewOrderHandler(dir string, orderService service.OrderService) *OrderHandler {
	repo := dal.NewRepository(dir)
	return &OrderHandler{orderService: orderService, repo: repo}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("Order created", order.ID)
	h.respondWithJSON(w, http.StatusCreated, order)
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.GetOrders()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.respondWithJSON(w, http.StatusOK, orders)
}

func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		h.respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
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

	log.Println("Order updated", order.ID)
	h.respondWithJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	if err := h.orderService.DeleteOrder(id); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("Order deleted", id)
	h.respondWithJSON(w, http.StatusOK, map[string]string{"message": "Order deleted successfully"})
}

func (h *OrderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing order ID")
		return
	}

	if err := h.orderService.CloseOrder(id); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("Order closed", id)
	h.respondWithJSON(w, http.StatusOK, map[string]string{"message": "Order closed successfully"})
}

func (h *OrderHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *OrderHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
