package main

import (
	"github.com/image-storage-service/cmd/api"
	"log"
)

func main() {
	log.Printf("Starting ImageStorageService...")

	api.StartServer()
}
