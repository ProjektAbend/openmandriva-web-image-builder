package logic

import (
	"encoding/json"
	"fmt"
	"github.com/shared/constants"
	"github.com/shared/models"
	"github.com/shared/status"
	"github.com/teris-io/shortid"
	"log"
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

	err = c.MessageBroker.CreateAndBindQueueToExchange(imageId, constants.EXCHANGE_STATUS, imageId)
	if err != nil {
		return "", err
	}

	c.BuildStatusHandler.SetStatusOfImageBuild(imageId, models.REQUESTED)

	jsonData, err := json.Marshal(imageConfig)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON %s", err)
	}

	if err := c.MessageBroker.SendMessageToQueue(string(jsonData), constants.QUEUE_BUILD); err != nil {
		return "", fmt.Errorf("error sending message to queue: %s", err)
	}
	c.BuildStatusHandler.SetStatusOfImageBuild(*imageConfig.ImageId, models.ACCEPTED)

	log.Printf("taking care of image %s", imageId)
	return imageId, nil
}

func (c *ImageBuilderLogic) generateImageId() (models.ImageId, error) {
	shortId, err := c.ShortIdGenerator.Generate()
	if err != nil {
		return "", err
	}
	return shortId, nil
}

func (c *ImageBuilderLogic) GetStatusOfImage(imageId models.ImageId) (models.ImageInfo, error) {
	messages, err := c.MessageBroker.CopyEveryMessageInsideStatusQueue(imageId)

	var imageBuildStatuses []status.ImageBuildStatus

	for _, message := range messages {
		var imageBuildStatus status.ImageBuildStatus
		err = json.Unmarshal(message, &imageBuildStatus)
		if err != nil {
			return models.ImageInfo{}, fmt.Errorf("error unmarshalling message: %s", err)
		}
		imageBuildStatuses = append(imageBuildStatuses, imageBuildStatus)
	}

	latestStatus := findLatestStatus(imageBuildStatuses)
	log.Printf("This is the latest status of image %s: %s", imageId, latestStatus)

	isAvailable := latestStatus == models.AVAILABLE

	imageInfo := models.ImageInfo{
		AvailableUntil: nil,
		ImageId:        imageId,
		Status:         &latestStatus,
		IsAvailable:    &isAvailable,
	}

	return imageInfo, nil
}

func findLatestStatus(imageBuildStatuses []status.ImageBuildStatus) models.ProcessingStatus {
	maxStatus := models.REQUESTED
	maxValue := 0

	for _, s := range imageBuildStatuses {
		if val, ok := status.Sequence[s.Status]; ok {
			if val > maxValue {
				maxValue = val
				maxStatus = s.Status
			}
		}
	}

	return maxStatus
}
