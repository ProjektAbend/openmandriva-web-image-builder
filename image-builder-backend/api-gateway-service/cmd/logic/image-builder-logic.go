package logic

import (
	"encoding/json"
	"fmt"
	"github.com/api-gateway-service/cmd/api"
)

type ImageBuilderLogic struct{}

func (c *ImageBuilderLogic) BuildImage(imageConfig api.ImageConfig) error {
	jsonData, err := json.Marshal(imageConfig)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	if err := sendMessageToQueue(string(jsonData), "buildQueue"); err != nil {
		return fmt.Errorf("error sending message to queue: %w", err)
	}

	return nil
}
