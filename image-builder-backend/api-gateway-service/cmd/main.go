package main

import (
	"github.com/api-gateway-service/cmd/api"
	"github.com/api-gateway-service/cmd/logic"
	"github.com/go-playground/validator/v10"
	"github.com/shared/constants"
	"github.com/shared/messagebroker"
	"github.com/teris-io/shortid"
)

func main() {
	messageBroker, _ := messagebroker.New(constants.LOCAL_HOST)
	shortIdGenerator, _ := shortid.New(1, shortid.DefaultABC, 2342)

	imageBuilderLogic := &logic.ImageBuilderLogic{
		MessageBroker:    messageBroker,
		ShortIdGenerator: shortIdGenerator,
	}

	validate := validator.New()

	api.StartServer(imageBuilderLogic, validate)
}
