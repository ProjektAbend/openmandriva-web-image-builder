package messagebroker

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shared/constants"
	"time"
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
		err := connection.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to close connection: %s", err)
		}
		return nil, fmt.Errorf("failed to open a channel: %s", err)
	}

	err = declareQueue(constants.BUILD_QUEUE, channel)
	if err != nil {
		return nil, err
	}

	return &MessageBroker{
		connection: connection,
		channel:    channel,
	}, nil
}

func declareQueue(name string, channel *amqp.Channel) error {
	_, err := channel.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error declaring queue: %s", err)
	}

	return nil
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
