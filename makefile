.PHONY: build
build:
	go build -o h ./cmd/hot-coffee/main.go

.DEFUALT_GOAL := build
