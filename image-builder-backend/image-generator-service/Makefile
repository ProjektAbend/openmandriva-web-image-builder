OPENAPI_FILE = ../../openapi.yml
TARGET_DIR = ./cmd/api

.PHONY: generate

all: generate

generate:
	oapi-codegen -generate types \
		-package api \
		$(OPENAPI_FILE) > $(TARGET_DIR)/types.gen.go
	go mod tidy
