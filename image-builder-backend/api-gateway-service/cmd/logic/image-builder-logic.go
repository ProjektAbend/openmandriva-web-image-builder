package logic

import (
	"encoding/json"
	"fmt"
	"github.com/api-gateway-service/cmd/api"
	"github.com/shared/constants"
	messagebroker "github.com/shared/message-broker"
	"github.com/teris-io/shortid"
)

type ImageBuilderLogic struct {
	MessageBroker    messagebroker.MessageBrokerInterface
	ShortIdGenerator *shortid.Shortid
}

func (c *ImageBuilderLogic) BuildImage(imageConfig api.ImageConfig) (api.ImageId, error) {
	imageId, err := c.generateImageId()
	if err != nil {
		return "", fmt.Errorf("error generating ImageId %s", err)
	}

	imageConfig.ImageId = &imageId

	jsonData, err := json.Marshal(imageConfig)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON %s", err)
	}

	if err := c.MessageBroker.SendMessageToQueue(string(jsonData), constants.BUILD_QUEUE); err != nil {
		return "", fmt.Errorf("error sending message to queue: %s", err)
	}
	return imageId, nil
}

func (c *ImageBuilderLogic) generateImageId() (api.ImageId, error) {
	shortId, err := c.ShortIdGenerator.Generate()
	if err != nil {
		return "", err
	}
	return shortId, nil
}
