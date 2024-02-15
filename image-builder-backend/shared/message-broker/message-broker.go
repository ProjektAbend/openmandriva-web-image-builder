package message_broker

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type MessageBrokerInterface interface {
	SendMessageToQueue(message string, queue string) error
	ConsumeMessage(queue string) (amqp.Delivery, error)
}

type MessageBroker struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func New() (*MessageBroker, error) {
	connection, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %s", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		connection.Close()
		return nil, fmt.Errorf("failed to open a channel: %s", err)
	}

	return &MessageBroker{
		connection: connection,
		channel:    channel,
	}, nil
}

func (c *MessageBroker) SendMessageToQueue(message string, queue string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.channel.PublishWithContext(ctx,
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %s", err)
	}

	return nil
}

func (c *MessageBroker) ConsumeMessage(queue string) (amqp.Delivery, error) {
	message, _, err := c.channel.Get(queue, true)
	if err != nil {
		return amqp.Delivery{}, err
	}

	return message, nil
}
