package logic

import (
	"encoding/json"
	"fmt"
	"github.com/api-gateway-service/cmd/api"
)

type ImageBuilderLogic struct{}

func (c *ImageBuilderLogic) BuildImage(imageConfig api.ImageConfig) (api.ImageInfo, error) {
	imageId, err := generateImageId()
	if err != nil {
		return api.ImageInfo{}, fmt.Errorf("error generating ImageId %s", err)
	}

	imageConfig.ImageId = &imageId
	jsonData, err := json.Marshal(imageConfig)
	if err != nil {
		return api.ImageInfo{}, fmt.Errorf("error marshalling JSON %s", err)
	}

	if err := sendMessageToQueue(string(jsonData), "buildQueue"); err != nil {
		return api.ImageInfo{}, fmt.Errorf("error sending message to queue: %s", err)
	}

	imageInfo := api.ImageInfo{
		ImageId: imageId,
	}

	return imageInfo, nil
}

func generateImageId() (api.ImageId, error) {
	// TODO: implement
	return "a1b2c3", nil
}
