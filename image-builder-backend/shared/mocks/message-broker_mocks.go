package mocks

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
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
			"architecture":"aarch64-uefi"
		}`),
	}, nil
}

func (_ *MockMessageBroker) CreateAndBindQueueToExchange(_ string, _ string, _ string) error {
	return nil
}

func (_ *MockMessageBroker) CopyEveryMessageInsideStatusQueue(_ string) ([][]byte, error) {
	return nil, nil
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
	return nil, nil
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
	return nil, nil
}
