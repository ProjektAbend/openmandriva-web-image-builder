package logic

import (
	"encoding/json"
	"fmt"
	"github.com/api-gateway-service/cmd/api"
)

type ImageBuilderLogic struct{}

func (c *ImageBuilderLogic) BuildImage(imageConfig api.ImageConfig) (api.ImageId, error) {
	imageId, err := generateImageId()
	if err != nil {
		return "", fmt.Errorf("error generating ImageId %s", err)
	}

	imageConfig.ImageId = &imageId

	jsonData, err := json.Marshal(imageConfig)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON %s", err)
	}

	if err := sendMessageToQueue(string(jsonData), "buildQueue"); err != nil {
		return "", fmt.Errorf("error sending message to queue: %s", err)
	}

	return imageId, nil
}

func generateImageId() (api.ImageId, error) {
	// TODO: implement
	// Keep track of every generated imageId on a list as long as the image with the
	// given imageId is available to download. Make sure the imageId does not exist yet
	// by comparing the generated imageId with the existing ones on the list.
	return "a1b2c3", nil
}
