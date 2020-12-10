test:
	go test -v ./...

build:
	go build -o ./rpggen.exe ./rpggen/main.go

.DEFAULT_GOAL := build
