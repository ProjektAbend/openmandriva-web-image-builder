package main

import (
	"github.com/image-generator-service/cmd/logic"
	"github.com/shared/message-broker"
	"log"
)

func main() {
	messageBroker, err := message_broker.New()
	if err != nil {
		log.Fatalf("Error trying to instantiate MessageBroker: %s", err)
	}

	generatorLogic := &logic.GeneratorLogic{
		MessageBroker: messageBroker,
	}

	generatorLogic.ProcessBuildRequests()

}
