package main

import (
	"github.com/image-generator-service/cmd/logic"
	"github.com/shared/messagebroker"
	"log"
)

func main() {
	log.Printf("Starting ImageGeneratorService...")

	messageBroker, err := messagebroker.New("rabbitmq")
	if err != nil {
		log.Fatalf("Error trying to instantiate MessageBroker: %s", err)
	}

	commandHandler := &logic.CommandHandler{}

	generatorLogic := &logic.GeneratorLogic{
		MessageBroker:  messageBroker,
		CommandHandler: commandHandler,
	}

	generatorLogic.ProcessBuildRequests()
}
