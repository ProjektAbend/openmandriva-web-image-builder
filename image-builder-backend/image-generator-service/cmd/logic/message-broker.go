package logic

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

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

func (c *MessageBroker) ConsumeMessage(queue string) (<-chan amqp.Delivery, error) {
	messages, err := c.channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
