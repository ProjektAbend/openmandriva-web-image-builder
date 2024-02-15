package logic

import (
	"encoding/json"
	"fmt"
	"github.com/image-generator-service/cmd/api"
	"github.com/shared/constants"
	"github.com/shared/message-broker"
	"log"
	"time"
)

type GeneratorLogic struct {
	MessageBroker message_broker.MessageBrokerInterface
}

func (c *GeneratorLogic) ProcessBuildRequests() {
	for {
		imageConfig, err := c.ProcessBuildRequest()
		if err != nil {
			continue
		}
		generateImage(imageConfig)
	}
}

func (c *GeneratorLogic) ProcessBuildRequest() (api.ImageConfig, error) {
	message, err := c.MessageBroker.ConsumeMessage(constants.BUILD_QUEUE)
	if err != nil {
		return api.ImageConfig{}, fmt.Errorf("error consuming message: %s", err)
	}

	var imageConfig api.ImageConfig
	err = json.Unmarshal(message.Body, &imageConfig)
	if err != nil {
		return api.ImageConfig{}, fmt.Errorf("error unmarshalling message: %s", err)
	}

	return imageConfig, nil
}

func generateImage(imageConfig api.ImageConfig) {
	log.Printf("Processing image with ID: %v", *imageConfig.ImageId)
	time.Sleep(5 * time.Second)
}
