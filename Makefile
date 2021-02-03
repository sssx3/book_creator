.DEFAULT_GOAL=build

build:
	go build -o bcreator ./cmd/bcreator/main.go 

run:
	go build -o bcreator ./cmd/bcreator/main.go
	./bcreator
