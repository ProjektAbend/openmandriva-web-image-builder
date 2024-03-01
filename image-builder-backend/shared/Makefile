OPENAPI_FILE = ../../openapi.yml
TARGET_DIR = ./models

.PHONY: generate

all: generate

generate:
	oapi-codegen -generate types \
		-package models \
		$(OPENAPI_FILE) > $(TARGET_DIR)/types.gen.go
	go mod tidy
