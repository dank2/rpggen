test:
	go test -v ./...

build:
	go build -o rpggen cmd/rpggen/main.go

.DEFAULT_GOAL := build
