package main

import (
	"github.com/api-gateway-service/cmd/api"
	"github.com/api-gateway-service/cmd/logic"
	"github.com/go-playground/validator/v10"
)

func main() {
	imageBuilderLogic := &logic.ImageBuilderLogic{}
	validate := validator.New()

	api.StartServer(imageBuilderLogic, validate)
}
