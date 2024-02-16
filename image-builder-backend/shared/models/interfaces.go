package models

import amqp "github.com/rabbitmq/amqp091-go"

type ImageBuilderLogicInterface interface {
	BuildImage(imageConfig ImageConfig) (ImageId, error)
}

type MessageBrokerInterface interface {
	SendMessageToQueue(message string, queue string) error
	ConsumeMessage(queue string) (amqp.Delivery, error)
}
