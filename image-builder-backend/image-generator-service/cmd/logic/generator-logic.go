package logic

import (
	"encoding/json"
	"fmt"
	"github.com/shared/constants"
	"github.com/shared/models"
	"log"
)

type GeneratorLogic struct {
	MessageBroker      models.MessageBrokerInterface
	BuildStatusHandler models.BuildStatusHandlerInterface
	CommandHandler     models.CommandHandlerInterface
}

func (c *GeneratorLogic) ProcessBuildRequests() {
	for {
		imageConfig, isEmpty, err := c.GetImageConfig()
		if err != nil {
			log.Printf("error while consuming message: %s", err)
		}
		if !isEmpty {
			c.BuildStatusHandler.SetStatusOfImageBuild(*imageConfig.ImageId, models.BUILDSTARTED)
			err = c.CommandHandler.GenerateImage(imageConfig)
			if err != nil {
				c.BuildStatusHandler.SetStatusOfImageBuild(*imageConfig.ImageId, models.BUILDFAILED)
				log.Printf("error while generating image: %s", err)
			}
			c.BuildStatusHandler.SetStatusOfImageBuild(*imageConfig.ImageId, models.BUILDFINISHED)
		}
	}
}

func (c *GeneratorLogic) GetImageConfig() (models.ImageConfig, bool, error) {
	message, err := c.MessageBroker.ConsumeMessage(constants.QUEUE_BUILD)
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

	c.BuildStatusHandler.SetStatusOfImageBuild(*imageConfig.ImageId, models.FETCHED)

	return imageConfig, false, nil
}
