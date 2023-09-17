OPENAPI_FILE = ../../openapi.yml
TARGET_DIR = ./cmd/api

.PHONY: generate

all: generate

generate:
	# generate types and server into different files
	# for better separation of concerns
	oapi-codegen -generate types \
		-package api \
		$(OPENAPI_FILE) > $(TARGET_DIR)/types.gen.go
	oapi-codegen -generate gin-server \
		-package api \
		$(OPENAPI_FILE) > $(TARGET_DIR)/interface.gen.go
	go mod tidy
