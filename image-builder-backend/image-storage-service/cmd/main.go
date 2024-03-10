package main

import (
	"github.com/image-storage-service/cmd/api"
	"github.com/image-storage-service/cmd/logic"
	"log"
)

func main() {
	log.Printf("Starting ImageStorageService...")

	imageStorageLogic := &logic.ImageStorageLogic{}
	api.StartServer(imageStorageLogic)
}
