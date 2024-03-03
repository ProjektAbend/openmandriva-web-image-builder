package status

import (
	"encoding/json"
	"github.com/shared/constants"
	"github.com/shared/models"
	"log"
)

var Sequence = map[models.Status]int{
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

type BuildStatusHandler struct {
	MessageBroker models.MessageBrokerInterface
}

func (c *BuildStatusHandler) SetStatusOfImageBuild(imageId models.ImageId, status models.Status) {
	log.Printf("set status of %s to %s", imageId, status)

	jsonData, err := json.Marshal(status)
	if err != nil {
		log.Printf("error marshalling JSON %s", err)
	}

	err = c.MessageBroker.SendMessageToExchange(string(jsonData), constants.EXCHANGE_STATUS, imageId)
	if err != nil {
		log.Printf("error sending message to exchange %s", err)
	}
}
