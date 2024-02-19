package logic

import (
	"encoding/json"
	"fmt"
	"github.com/shared/constants"
	"github.com/shared/models"
	"log"
)

type GeneratorLogic struct {
	MessageBroker  models.MessageBrokerInterface
	CommandHandler models.CommandHandlerInterface
}

func (c *GeneratorLogic) ProcessBuildRequests() {
	for {
		imageConfig, isEmpty, err := c.ProcessBuildRequest()
		if err != nil {
			log.Printf("error while consuming message: %s", err)
		}
		if !isEmpty {
			c.generateImage(imageConfig)
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

func (c *GeneratorLogic) generateImage(imageConfig models.ImageConfig) {
	log.Printf("Processing image with ID: %v", *imageConfig.ImageId)
	err := c.CommandHandler.RunCommand("./os-image-builder/build", imageConfig.Architecture)
	if err != nil {
		log.Printf("error running command: %s", err)
	}
}
