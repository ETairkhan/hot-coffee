package server

import (
	"ayzhunis/hot-coffee/internal/handler"
	"errors"
	"fmt"
	"net/http"
)

var ()

type server struct {
	port int
	Dir  string

	handler handler.Handler

	mux *http.ServeMux
}

func NewServer(port int, dir string) (*server, error) {
	if port <= 0 || port >= 63535 {
		return nil, errors.New("invalid port")
	}

	s := server{
		port: port,
		Dir:  dir,
		mux:  http.NewServeMux(),
	}

	return &s, nil
}

func (s *server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}
