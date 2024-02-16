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
		imageConfig, isEmpty, err := c.ProcessBuildRequest()
		if err != nil {
			log.Printf("error while consuming message: %s", err)
		}
		if !isEmpty {
			generateImage(imageConfig)
		}
	}
}

func (c *GeneratorLogic) ProcessBuildRequest() (models.ImageConfig, bool, error) {
	message, err := c.MessageBroker.ConsumeMessage(constants.BUILD_QUEUE)
	if err != nil {
		return models.ImageConfig{}, false, fmt.Errorf("error consuming message: %s", err)
	}

	if message.Body == nil {
		return models.ImageConfig{}, true, nil
	}

	var imageConfig models.ImageConfig
	err = json.Unmarshal(message.Body, &imageConfig)
	if err != nil {
		return models.ImageConfig{}, false, fmt.Errorf("error unmarshalling message: %s", err)
	}

	return imageConfig, false, nil
}

func generateImage(imageConfig models.ImageConfig) {
	log.Printf("Processing image with ID: %v", *imageConfig.ImageId)
	time.Sleep(5 * time.Second)
}
