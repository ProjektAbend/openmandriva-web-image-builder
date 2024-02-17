package main

import (
	"github.com/image-generator-service/cmd/logic"
	"github.com/shared/messagebroker"
	"log"
)

func main() {
	log.Printf("Starting ImageGeneratorService...")

	messageBroker, err := messagebroker.New()
	if err != nil {
		log.Fatalf("Error trying to instantiate MessageBroker: %s", err)
	}

	generatorLogic := &logic.GeneratorLogic{
		MessageBroker: messageBroker,
	}

	generatorLogic.ProcessBuildRequests()
}
