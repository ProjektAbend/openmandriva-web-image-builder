package logic

import (
	"encoding/json"
	"fmt"
	"github.com/shared/constants"
	"github.com/shared/models"
	"github.com/teris-io/shortid"
	"log"
	"math"
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

	err = c.MessageBroker.CreateAndBindQueueToExchange(imageId, constants.EXCHANGE_STATUS, imageId)
	if err != nil {
		return "", err
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
	messages, err := c.MessageBroker.ConsumeMessages(imageId)
	if err != nil {
		return models.ImageInfo{}, fmt.Errorf("error consuming messages: %s", err)
	}

	var imageBuildStatuses []models.ImageBuildStatus

	for _, m := range messages {
		var imageBuildStatus models.ImageBuildStatus
		err = json.Unmarshal(m, &imageBuildStatus)
		if err != nil {
			return models.ImageInfo{}, fmt.Errorf("error unmarshalling message: %s", err)
		}
		imageBuildStatuses = append(imageBuildStatuses, imageBuildStatus)
	}

	// TODO: integrate latestStatus into imageInfo for a more fine grained status
	latestStatus := findLatestStatus(imageBuildStatuses)
	log.Printf("This is the latest status of image %s: %s", imageId, latestStatus)

	isAvailable := latestStatus == models.AVAILABLE

	imageInfo := models.ImageInfo{
		AvailableUntil: nil,
		ImageId:        imageId,
		IsAvailable:    &isAvailable,
	}

	return imageInfo, nil
}

func findLatestStatus(statuses []models.ImageBuildStatus) models.ImageProcessingStatus {
	latestStatus := models.ImageProcessingStatus(math.MinInt32)
	for _, status := range statuses {
		if status.Status > latestStatus {
			latestStatus = status.Status
		}
	}
	return latestStatus
}
