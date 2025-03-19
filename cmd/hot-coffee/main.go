package main

import (
	"flag"
	"fmt"
	"log/slog"

	"ayzhunis/hot-coffee/internal/server"
)

var (
	port int
	dir  string
)

func init() {
	flag.IntVar(&port, "port", 8080, "port number.")
	flag.StringVar(&dir, "dir", "data", "Path to the data directory")
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	s, err := server.NewServer(port, dir)
	if err != nil {
		return
	}
	if err := s.Run(); err != nil {
		slog.Error(err.Error())
		return
	}
}

func Usage() {
	fmt.Println(`$ ./hot-coffee --help
Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.`)
}
