package server

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/internal/handler"
	"ayzhunis/hot-coffee/internal/service"
	"errors"
	"fmt"
	"net/http"
)

var ()

type server struct {
	port int
	Dir  string

	handler *handler.OrderHandler

	mux *http.ServeMux
}

func NewServer(port int, dir string) (*server, error) {
	if port <= 0 || port >= 63535 {
		return nil, errors.New("invalid port")
	}
	orderRepository := dal.NewOrderRepository(dir)
	serv := service.NewOrderService(orderRepository)
	handler := handler.NewOrderHandler(serv)

	s := server{
		port:    port,
		Dir:     dir,
		mux:     http.NewServeMux(),
		handler: handler,
	}

	s.registerRoutes()

	return &s, nil
}

// registerRoutes sets up HTTP routes for order handling
func (s *server) registerRoutes() {
	s.mux.HandleFunc("GET /orders", s.handler.GetOrders) // GET all orders
	s.mux.HandleFunc("GET /orders/{id}", s.handler.GetOrderByID) // GET order by id
	s.mux.HandleFunc("POST /orders", s.handler.CreateOrder) // POST create order
	// s.mux.HandleFunc("PUT/orders/update", s.handler.UpdateOrder) // PUT update order
	// s.mux.HandleFunc("DELETE /orders/delete", s.handler.DeleteOrder) // DELETE delete order
	// s.mux.HandleFunc("POST /orders/close", s.handler.CloseOrder)   // POST close order
}

func (s *server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}
