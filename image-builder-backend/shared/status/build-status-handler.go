package status

import (
	"encoding/json"
	"github.com/shared/constants"
	"github.com/shared/models"
	"log"
)

type BuildStatusHandler struct {
	MessageBroker models.MessageBrokerInterface
}

func (c *BuildStatusHandler) SetStatusOfImageBuild(imageId models.ImageId, status models.ImageProcessingStatus) {
	log.Printf("set status of %s to %s", imageId, status)
	imageBuildStatus := &models.ImageBuildStatus{
		ImageId: imageId,
		Status:  status,
	}

	jsonData, err := json.Marshal(imageBuildStatus)
	if err != nil {
		log.Printf("error marshalling JSON %s", err)
	}

	_ = c.MessageBroker.SendMessageToExchange(string(jsonData), constants.EXCHANGE_STATUS, imageId)
}
