package server

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"ayzhunis/hot-coffee/internal/dal"
	"ayzhunis/hot-coffee/internal/handler"
	"ayzhunis/hot-coffee/internal/service"
	"ayzhunis/hot-coffee/utils"

)

type server struct {
	port int
	Dir  string

	orderHandler       *handler.OrderHandler
	menuHandler        *handler.MenuHandler
	inventoryHandler   *handler.InventoryHandler
	aggregationHandler *handler.AggregationHandler

	mux *http.ServeMux
}

func NewServer(port int, dir string) (*server, error) {
	if port <= 0 || port >= 63535 {
		return nil, errors.New("invalid port")
	}

	if err := utils.CheckDir(dir); err != nil {
		return nil, err
	}

	orderRepository := dal.NewOrderRepository(dir)
	menuRepository := dal.NewMenuRepository(dir)
	inventoryRepository := dal.NewInventoryRepository(dir)

	// Service
	menuServ := service.NewMenuService(menuRepository, inventoryRepository)
	inventoryServ := service.NewInventoryService(inventoryRepository)
	orderServ := service.NewOrderService(orderRepository, menuServ, inventoryServ)
	aggregationServ := service.NewAggregationService(orderRepository, menuRepository, inventoryRepository)

	// Handler
	orderHandler := handler.NewOrderHandler(orderServ)
	menuHandler := handler.NewMenuHandler(menuServ)
	inventoryHandler := handler.NewInventoryHandler(inventoryServ)
	aggregationHandler := handler.NewAggregationHandler(aggregationServ)

	s := server{
		port:               port,
		Dir:                dir,
		mux:                http.NewServeMux(),
		orderHandler:       orderHandler,
		menuHandler:        menuHandler,
		inventoryHandler:   inventoryHandler,
		aggregationHandler: aggregationHandler,
	}

	s.registerRoutes()

	return &s, nil
}

// registerRoutes sets up HTTP routes for order handling
func (s *server) registerRoutes() {
	s.mux.HandleFunc("POST /orders", s.orderHandler.CreateOrder)
	s.mux.HandleFunc("GET /orders", s.orderHandler.GetOrders)
	s.mux.HandleFunc("GET /orders/{id}", s.orderHandler.GetOrderByID)
	s.mux.HandleFunc("PUT /orders/{id}", s.orderHandler.UpdateOrder)
	s.mux.HandleFunc("DELETE /orders/{id}", s.orderHandler.DeleteOrder)
	s.mux.HandleFunc("POST /orders/{id}/close", s.orderHandler.CloseOrder)

	s.mux.HandleFunc("POST /menu", s.menuHandler.CreateMenu)
	s.mux.HandleFunc("GET /menu", s.menuHandler.GetAllMenuItems)
	s.mux.HandleFunc("GET /menu/{id}", s.menuHandler.GetMenuItemByID)
	s.mux.HandleFunc("PUT /menu/{id}", s.menuHandler.UpdateMenuItem)
	s.mux.HandleFunc("DELETE /menu/{id}", s.menuHandler.DeleteMenuItemById)

	s.mux.HandleFunc("POST /inventory", s.inventoryHandler.CreateInventoryItems)
	s.mux.HandleFunc("GET /inventory", s.inventoryHandler.GetAllInventory)
	s.mux.HandleFunc("GET /inventory/{id}", s.inventoryHandler.GetInventoryById)
	s.mux.HandleFunc("PUT /inventory/{id}", s.inventoryHandler.UpdateInventoryItem)
	s.mux.HandleFunc("DELETE /inventory/{id}", s.inventoryHandler.DeleteInventoryItem)

	s.mux.HandleFunc("GET /reports/total-sales", s.aggregationHandler.TotalSales)
	s.mux.HandleFunc("GET /reports/popular-items", s.aggregationHandler.PopularItems)
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
