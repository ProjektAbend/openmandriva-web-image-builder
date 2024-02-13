package logic

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type GeneratorLogic struct {
	MessageBroker *MessageBroker
}

func (c *GeneratorLogic) ProcessBuildRequests() {
	for {
		message, err := c.MessageBroker.ConsumeMessage("buildQueue")
		if err != nil {
			log.Printf("Error consuming message: %s", err)
			return
		}

		generateImage(message)
	}
}

func generateImage(message amqp.Delivery) {
	log.Printf("Received a message: %s", message.Body)
	time.Sleep(10 * time.Second)
}
