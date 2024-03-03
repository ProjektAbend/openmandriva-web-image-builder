package status

import (
	"encoding/json"
	"github.com/shared/constants"
	"github.com/shared/models"
	"log"
)

var Sequence = map[models.ProcessingStatus]int{
	models.REQUESTED:     1,
	models.ACCEPTED:      2,
	models.FETCHED:       3,
	models.BUILDSTARTED:  4,
	models.BUILDFAILED:   5,
	models.BUILDFINISHED: 6,
	models.UPLOADSTARTED: 7,
	models.UPLOADFAILED:  8,
	models.AVAILABLE:     9,
	models.EXPIRED:       10,
}

type ImageBuildStatus struct {
	ImageId models.ImageId
	Status  models.ProcessingStatus
}

type BuildStatusHandler struct {
	MessageBroker models.MessageBrokerInterface
}

func (c *BuildStatusHandler) SetStatusOfImageBuild(imageId models.ImageId, status models.ProcessingStatus) {
	log.Printf("set status of %s to %s", imageId, status)
	imageBuildStatus := &ImageBuildStatus{
		ImageId: imageId,
		Status:  status,
	}

	jsonData, err := json.Marshal(imageBuildStatus)
	if err != nil {
		log.Printf("error marshalling JSON %s", err)
	}

	err = c.MessageBroker.SendMessageToExchange(string(jsonData), constants.EXCHANGE_STATUS, imageId)
	if err != nil {
		log.Printf("error sending message to exchange %s", err)
	}
}
