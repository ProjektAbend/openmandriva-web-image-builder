package main

import (
	"github.com/api-gateway-service/cmd/api"
	"github.com/api-gateway-service/cmd/logic"
	"github.com/go-playground/validator/v10"
	"github.com/shared/constants"
	"github.com/shared/messagebroker"
	"github.com/shared/status"
	"github.com/teris-io/shortid"
	"log"
)

func main() {
	messageBroker, err := messagebroker.New(constants.LOCAL_HOST)
	if err != nil {
		log.Fatalf("Error trying to instantiate MessageBroker: %s", err)
	}

	buildStatusHandler := &status.BuildStatusHandler{
		MessageBroker: messageBroker,
	}

	shortIdGenerator, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		log.Fatalf("Error trying to instantiate ShortIdGenerator: %s", err)
	}

	imageBuilderLogic := &logic.ImageBuilderLogic{
		MessageBroker:      messageBroker,
		BuildStatusHandler: buildStatusHandler,
		ShortIdGenerator:   shortIdGenerator,
	}

	validate := validator.New()

	api.StartServer(imageBuilderLogic, validate)
}
