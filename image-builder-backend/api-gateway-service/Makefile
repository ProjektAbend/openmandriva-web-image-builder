.PHONY: generate build run

all: build run

generate:
	# reuse openapi-generator-cli from frontend
	../../image-builder-frontend/node_modules/.bin/openapi-generator-cli generate -i ../../openapi.yml -g go-gin-server -o ./api

build:
	go mod tidy
	go build -o cmd ./api/main.go

run:
	./api-gateway-service
