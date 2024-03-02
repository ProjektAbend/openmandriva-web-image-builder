package logic

import (
	"encoding/json"
	"fmt"
	"github.com/shared/constants"
	"github.com/shared/models"
	"github.com/teris-io/shortid"
)

type ImageBuilderLogic struct {
	MessageBroker      models.MessageBrokerInterface
	BuildStatusHandler models.BuildStatusHandlerInterface
	ShortIdGenerator   *shortid.Shortid
}

func (c *ImageBuilderLogic) BuildImage(imageConfig models.ImageConfig) (models.ImageId, error) {
	imageId, err := c.generateImageId()
	if err != nil {
		return "", fmt.Errorf("error generating ImageId %s", err)
	}

	imageConfig.ImageId = &imageId

	c.BuildStatusHandler.SetStatusOfImageBuild(imageId, models.REQUESTED)

	jsonData, err := json.Marshal(imageConfig)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON %s", err)
	}

	if err := c.MessageBroker.SendMessageToQueue(string(jsonData), constants.QUEUE_BUILD); err != nil {
		return "", fmt.Errorf("error sending message to queue: %s", err)
	}
	c.BuildStatusHandler.SetStatusOfImageBuild(*imageConfig.ImageId, models.ACCEPTED)
	return imageId, nil
}

func (c *ImageBuilderLogic) generateImageId() (models.ImageId, error) {
	shortId, err := c.ShortIdGenerator.Generate()
	if err != nil {
		return "", err
	}
	return shortId, nil
}
