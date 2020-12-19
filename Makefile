test:
	go test ./... --ginkgo.noColor

build:
	go build -o ./bin/rpggen.exe ./rpggen/main.go

install:
	go build -o ./bin/rpggen.exe ./rpggen/main.go \
	&& mv ./rpggen.exe ~/Tools/rpggen.exe

.DEFAULT_GOAL := build
