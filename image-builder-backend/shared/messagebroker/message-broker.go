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

func New(host string) (*MessageBroker, error) {
	connection, err := amqp.Dial("amqp://admin:admin@" + host + ":5672/")
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

	err = declareQueue(constants.QUEUE_BUILD, channel)
	if err != nil {
		return nil, err
	}

	err = declareExchange(constants.EXCHANGE_STATUS, "direct", channel)
	if err != nil {
		return nil, err
	}

	return &MessageBroker{
		connection: connection,
		channel:    channel,
	}, nil
}

func (c *MessageBroker) SendMessageToQueue(message string, queueName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.channel.PublishWithContext(ctx,
		"",
		queueName,
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

func (c *MessageBroker) SendMessageToExchange(message string, exchangeName string, routingKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.channel.PublishWithContext(ctx,
		exchangeName,
		routingKey,
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

func (c *MessageBroker) CopyEveryMessageInsideStatusQueue(queue string) ([][]byte, error) {
	channel, _ := c.connection.Channel()
	var messages [][]byte
	for {
		message, err := consumeMessageNoAck(queue, channel)
		if err != nil {
			return nil, fmt.Errorf("error consuming message from status queue: %s", err)
		}
		if message.Body != nil {
			messages = append(messages, message.Body)
		} else {
			break
		}
	}
	channel.Close()
	return messages, nil
}

func consumeMessageNoAck(queue string, channel *amqp.Channel) (amqp.Delivery, error) {
	message, _, err := channel.Get(queue, false)
	if err != nil {
		return amqp.Delivery{}, err
	}

	return message, nil
}

func (c *MessageBroker) CreateAndBindQueueToExchange(queueName string, exchangeName string, routingKey string) error {
	err := declareQueue(queueName, c.channel)
	if err != nil {
		return err
	}

	err = c.channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error binding queue %s to exchange %s: %s", queueName, exchangeName, err)
	}

	return nil
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

func declareExchange(name string, exchangeType string, channel *amqp.Channel) error {
	err := channel.ExchangeDeclare(
		name,
		exchangeType,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error declaring exchange: %s", err)
	}
	return nil
}
