package main

import (
	"github.com/image-generator-service/cmd/logic"
	"github.com/shared/constants"
	"github.com/shared/messagebroker"
	"github.com/shared/mocks"
	"github.com/shared/models"
	"github.com/shared/status"
	"log"
	"os"
)

func main() {
	log.Printf("Starting ImageGeneratorService...")

	messageBroker, err := messagebroker.New(getMessageBrokerHost())
	if err != nil {
		log.Fatalf("Error trying to instantiate MessageBroker: %s", err)
	}

	buildStatusHandler := &status.BuildStatusHandler{
		MessageBroker: messageBroker,
	}

	var commandHandler models.CommandHandlerInterface
	if useMocks() {
		log.Printf("Run ImageGeneratorService in MOCK MODE")
		commandHandler = &mocks.MockCommandHandlerGenerateFakeIso{}
	} else {
		commandHandler = &logic.CommandHandler{}
	}

	generatorLogic := &logic.GeneratorLogic{
		MessageBroker:      messageBroker,
		BuildStatusHandler: buildStatusHandler,
		CommandHandler:     commandHandler,
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
			return true
		}
	}
	return false
}
