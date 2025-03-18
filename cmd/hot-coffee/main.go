package main

import "flag"

var (
	port int
	dir string
)

func init() {
	flag.IntVar(&port, "port", 8080, "port number.")
	flag.StringVar(&dir, "dir", "data", "Path to the data directory")
}

func main() {
	flag.Parse()
	
	
}