package main

import (
	"github.com/image-generator-service/cmd/logic"
	"github.com/shared/constants"
	"github.com/shared/messagebroker"
	"github.com/shared/mocks"
	"github.com/shared/models"
	"log"
	"os"
)

func main() {
	log.Printf("Starting ImageGeneratorService...")

	messageBroker, err := messagebroker.New(getMessageBrokerHost())
	if err != nil {
		log.Fatalf("Error trying to instantiate MessageBroker: %s", err)
	}

	var commandHandler models.CommandHandlerInterface
	if useMocks() {
		commandHandler = &mocks.MockCommandHandlerGenerateFakeIso{}
	} else {
		commandHandler = &logic.CommandHandler{}
	}

	generatorLogic := &logic.GeneratorLogic{
		MessageBroker:  messageBroker,
		CommandHandler: commandHandler,
	}

	generatorLogic.ProcessBuildRequests()
}

func getMessageBrokerHost() string {
	envVariable := os.Getenv("RABBIT_MQ_HOST")
	if envVariable == "" {
		return constants.LOCAL_HOST
	}
	return envVariable
}

func useMocks() bool {
	for _, arg := range os.Args {
		if arg == "--mock" {
			log.Printf("Run ImageGeneratorService in MOCK MODE")
			return true
		}
	}
	return false
}
