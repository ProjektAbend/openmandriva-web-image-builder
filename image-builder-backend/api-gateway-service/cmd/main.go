package main

import (
	"github.com/api-gateway-service/cmd/api"
	"github.com/api-gateway-service/cmd/logic"
)

func main() {
	imageBuilderLogic := &logic.ImageBuilderLogic{}

	api.StartServer(imageBuilderLogic)
}
