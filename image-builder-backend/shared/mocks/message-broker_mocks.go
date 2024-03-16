package mocks

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shared/models"
)

type MockMessageBroker struct{}

func (_ *MockMessageBroker) SendMessageToQueue(_ string, _ string) error {
	return nil
}

func (_ *MockMessageBroker) SendMessageToExchange(_ string, _ string, _ string) error {
	return nil
}

func (_ *MockMessageBroker) ConsumeMessage(_ string) (amqp.Delivery, error) {
	return amqp.Delivery{
		Body: []byte(`{
			"architecture":"aarch64-uefi",
			"imageId":"WZ3h633-p"
		}`),
	}, nil
}

func (_ *MockMessageBroker) CreateAndBindQueueToExchange(_ string, _ string, _ string) error {
	return nil
}

func (_ *MockMessageBroker) CopyEveryMessageInsideStatusQueue(_ string) ([][]byte, error) {
	var messages [][]byte
	message1, _ := json.Marshal(models.REQUESTED)
	message2, _ := json.Marshal(models.ACCEPTED)
	messages = append(messages, message1)
	messages = append(messages, message2)
	return messages, nil
}

type MockMessageBrokerReturnsError struct{}

func (_ *MockMessageBrokerReturnsError) SendMessageToQueue(_ string, _ string) error {
	return fmt.Errorf("error occurred")
}

func (_ *MockMessageBrokerReturnsError) SendMessageToExchange(_ string, _ string, _ string) error {
	return nil
}

func (_ *MockMessageBrokerReturnsError) ConsumeMessage(_ string) (amqp.Delivery, error) {
	return amqp.Delivery{}, fmt.Errorf("error occurred")
}

func (_ *MockMessageBrokerReturnsError) CreateAndBindQueueToExchange(_ string, _ string, _ string) error {
	return nil
}

func (_ *MockMessageBrokerReturnsError) CopyEveryMessageInsideStatusQueue(_ string) ([][]byte, error) {
	return nil, fmt.Errorf("error occurred")
}

type MockMessageBrokerHasEmptyQueue struct{}

func (_ *MockMessageBrokerHasEmptyQueue) SendMessageToQueue(_ string, _ string) error {
	return nil
}

func (_ *MockMessageBrokerHasEmptyQueue) SendMessageToExchange(_ string, _ string, _ string) error {
	return nil
}

func (_ *MockMessageBrokerHasEmptyQueue) ConsumeMessage(_ string) (amqp.Delivery, error) {
	return amqp.Delivery{}, nil
}

func (_ *MockMessageBrokerHasEmptyQueue) CreateAndBindQueueToExchange(_ string, _ string, _ string) error {
	return nil
}

func (_ *MockMessageBrokerHasEmptyQueue) CopyEveryMessageInsideStatusQueue(_ string) ([][]byte, error) {
	return [][]byte{}, nil
}
