OPENAPI_FILE = ../../openapi.yml
TARGET_DIR = ./cmd/api

IMAGE_STORAGE_OPENAPI_FILE = ../image-storage-service/openapi.yml

.PHONY: generate

all: generate

generate:
	# generate code for server
	oapi-codegen -generate gin-server,types \
		-package api \
		$(OPENAPI_FILE) > $(TARGET_DIR)/interface.gen.go

	#generate code for image-storage-client
	oapi-codegen -generate client,types \
		-package logic \
		$(IMAGE_STORAGE_OPENAPI_FILE) > ./cmd/logic/image-storage-client.gen.go

	go mod tidy