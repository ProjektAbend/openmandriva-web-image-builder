package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shared/constants"
	"github.com/shared/models"
	"github.com/shared/status"
	"github.com/teris-io/shortid"
	"io"
	"log"
	"net/http"
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

func (c *ImageBuilderLogic) GetImage(imageId models.ImageId) ([]byte, error) {
	client, err := NewClient("http://localhost:8084")
	if err != nil {
		return nil, fmt.Errorf("error creating image-storage client: %s", err)
	}

	fileName := imageId + ".iso"
	response, err := client.GetIsoFile(context.Background(), fileName)
	if err != nil {
		return nil, fmt.Errorf("error while calling GetIsoFile: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	fileContent, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

func (c *ImageBuilderLogic) GetStatusOfImage(imageId models.ImageId) (models.ImageInfo, error) {
	messages, err := c.MessageBroker.CopyEveryMessageInsideStatusQueue(imageId)
	if err != nil {
		return models.ImageInfo{}, fmt.Errorf("error copying messages from status queue: %s", err)
	}

	processingStatuses, err := unmarshalMessages(messages)
	if err != nil {
		return models.ImageInfo{}, fmt.Errorf("error unmarshalling message: %s", err)
	}

	latestStatus := FindLatestStatus(processingStatuses)
	log.Printf("This is the latest status of image %s: %s", imageId, latestStatus)

	imageInfo := models.ImageInfo{
		AvailableUntil: nil,
		ImageId:        imageId,
		Status:         &latestStatus,
	}

	return imageInfo, nil
}

func unmarshalMessages(messages [][]byte) ([]models.Status, error) {
	var processingStatuses []models.Status
	for _, message := range messages {
		var processingStatus models.Status
		err := json.Unmarshal(message, &processingStatus)
		if err != nil {
			return nil, err
		}
		processingStatuses = append(processingStatuses, processingStatus)
	}
	return processingStatuses, nil
}

func FindLatestStatus(processingStatuses []models.Status) models.Status {
	if len(processingStatuses) == 0 {
		return models.DOESNOTEXIST
	}

	maxStatus := models.REQUESTED
	maxValue := 0
	for _, s := range processingStatuses {
		if val, ok := status.Sequence[s]; ok {
			if val > maxValue {
				maxValue = val
				maxStatus = s
			}
		}
	}
	return maxStatus
}
