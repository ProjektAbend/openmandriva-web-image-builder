.PHONY: generate

all: generate

generate:
	mkdir api
	oapi-codegen -package api -generate types,gin-server ../../openapi.yml > ./api/api.gen.go
	go mod tidy
