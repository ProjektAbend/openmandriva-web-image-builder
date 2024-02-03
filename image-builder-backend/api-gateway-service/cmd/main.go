package main

import (
	"github.com/api-gateway-service/cmd/api"
	"github.com/api-gateway-service/cmd/logic"
	"github.com/go-playground/validator/v10"
)

func main() {
	messageBrokerLogic := &logic.MessageBroker{}
	imageBuilderLogic := &logic.ImageBuilderLogic{
		MessageBroker: messageBrokerLogic,
	}
	validate := validator.New()

	api.StartServer(imageBuilderLogic, validate)
}
