OPENAPI_FILE = ../../openapi.yml
TARGET_DIR = ./cmd/api

.PHONY: generate

all: generate

generate:
	oapi-codegen -generate gin-server,types \
		-package api \
		$(OPENAPI_FILE) > $(TARGET_DIR)/interface.gen.go
	go mod tidy