package main

import (
	"github.com/image-generator-service/cmd/logic"
	"log"
)

func main() {
	messageBroker, err := logic.New()
	if err != nil {
		log.Fatalf("Error trying to instantiate MessageBroker: %s", err)
	}

	generatorLogic := &logic.GeneratorLogic{
		MessageBroker: messageBroker,
	}

	generatorLogic.ProcessBuildRequests()

}
