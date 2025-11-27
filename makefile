.PHONY: default run build clean


default: run

run:
	go run main.go

build:
	go build -o bin/gestao_estoque_api main.go

docs:
	swag init -g main.go -o docs/

clean:
	-rmdir /s /q bin
	-rmdir /s /q docs