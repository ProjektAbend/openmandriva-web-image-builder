package logic

import (
	"encoding/json"
	"fmt"
	"github.com/shared/constants"
	"github.com/shared/models"
	"log"
	"time"
)

type GeneratorLogic struct {
	MessageBroker models.MessageBrokerInterface
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

func (c *GeneratorLogic) ProcessBuildRequest() (models.ImageConfig, error) {
	message, err := c.MessageBroker.ConsumeMessage(constants.BUILD_QUEUE)
	if err != nil {
		return models.ImageConfig{}, fmt.Errorf("error consuming message: %s", err)
	}

	var imageConfig models.ImageConfig
	err = json.Unmarshal(message.Body, &imageConfig)
	if err != nil {
		return models.ImageConfig{}, fmt.Errorf("error unmarshalling message: %s", err)
	}

	return imageConfig, nil
}

func generateImage(imageConfig models.ImageConfig) {
	log.Printf("Processing image with ID: %v", *imageConfig.ImageId)
	time.Sleep(5 * time.Second)
}
