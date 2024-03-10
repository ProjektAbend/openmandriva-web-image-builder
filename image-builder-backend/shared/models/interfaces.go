package models

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"mime/multipart"
)

type ImageBuilderLogicInterface interface {
	BuildImage(imageConfig ImageConfig) (ImageId, error)
	GetStatusOfImage(imageId ImageId) (ImageInfo, error)
}

type ImageStorageLogicInterface interface {
	StoreImage(file multipart.File, filename string) error
}

type MessageBrokerInterface interface {
	SendMessageToQueue(message string, queue string) error
	SendMessageToExchange(message string, exchangeName string, routingKey string) error
	CreateAndBindQueueToExchange(queueName string, exchangeName string, routingKey string) error
	CopyEveryMessageInsideStatusQueue(queue string) ([][]byte, error)
	ConsumeMessage(queue string) (amqp.Delivery, error)
}

type CommandHandlerInterface interface {
	GenerateImage(imageConfig ImageConfig) error
}

type BuildStatusHandlerInterface interface {
	SetStatusOfImageBuild(imageId ImageId, status Status)
}
