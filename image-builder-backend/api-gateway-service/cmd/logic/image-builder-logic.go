package logic

import (
	"encoding/json"
	"fmt"
	"github.com/api-gateway-service/cmd/api"
)

type ImageBuilderLogic struct{}

func (c *ImageBuilderLogic) BuildImage(imageConfig api.ImageConfig) error {
	obj := imageConfig

	jsonData, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	sendMessageToQueue(string(jsonData), "buildQueue")

	return nil
}
