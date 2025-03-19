package server

import (
	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/internal/handler"
	"ayzhunis/hot-coffee/internal/service"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type server struct {
	port int
	Dir  string

	orderHandler *handler.OrderHandler
	menuHandler *handler.MenuHandler

	mux *http.ServeMux
}

func NewServer(port int, dir string) (*server, error) {
	if port <= 0 || port >= 63535 {
		return nil, errors.New("invalid port")
	}
	orderRepository := dal.NewOrderRepository(dir)
	menuRepository := dal.NewMenuRepository(dir)

	orderServ := service.NewOrderService(orderRepository)
	menuServ := service.NewMenuService(menuRepository)

	orderHandler := handler.NewOrderHandler(orderServ)
	menuHandler := handler.NewMenuHandler(menuServ)

	s := server{
		port:    port,
		Dir:     dir,
		mux:     http.NewServeMux(),
		orderHandler: orderHandler,
		menuHandler: menuHandler,
	}

	s.registerRoutes()

	return &s, nil
}

// registerRoutes sets up HTTP routes for order handling
func (s *server) registerRoutes() {
	s.mux.HandleFunc("POST /orders", s.orderHandler.CreateOrder)      // POST create order
	s.mux.HandleFunc("GET /orders", s.orderHandler.GetOrders)         // GET all orders
	s.mux.HandleFunc("GET /orders/{id}", s.orderHandler.GetOrderByID) // GET order by id
	s.mux.HandleFunc("PUT /orders/{id}", s.orderHandler.UpdateOrder) // PUT update order
	s.mux.HandleFunc("DELETE /orders/{id}", s.orderHandler.DeleteOrder) // DELETE delete order
	s.mux.HandleFunc("POST /orders/{id}/close", s.orderHandler.CloseOrder)   // POST close order

	s.mux.HandleFunc("POST /menu", s.menuHandler.CreateMenu)
	s.mux.HandleFunc("GET /menu", s.menuHandler.GetAllMenuItems)
	s.mux.HandleFunc("GET /menu/{id}", s.menuHandler.GetMenuItemByID)
	s.mux.HandleFunc("PUT /menu/{id}", s.menuHandler.UpdateMenuItem)
	s.mux.HandleFunc("DELETE /menu/{id}", s.menuHandler.DeleteMenuItemById)
}

func (s *server) Run() error {
	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, handlerOpts))
	slog.SetDefault(logger)

	log.Printf("Starting the server on %d...\n", s.port)
	log.Printf("Data dir: %s", s.Dir)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}
