.PHONY: build

all: build

build:
	# build executable for linux in order to run it inside a docker container
	GOOS=linux GOARCH=amd64 go build -o build/image-generator-service cmd/main.go