package mocks

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MockMessageBroker struct{}

func (_ *MockMessageBroker) SendMessageToQueue(_ string, _ string) error {
	return nil
}

func (_ *MockMessageBroker) ConsumeMessage(_ string) (amqp.Delivery, error) {
	return amqp.Delivery{
		Body: []byte(`{
			"architecture":"aarch64-uefi"
		}`),
	}, nil
}

type MockMessageBrokerReturnsError struct{}

func (_ *MockMessageBrokerReturnsError) SendMessageToQueue(_ string, _ string) error {
	return fmt.Errorf("error occurred")
}

func (_ *MockMessageBrokerReturnsError) ConsumeMessage(_ string) (amqp.Delivery, error) {
	return amqp.Delivery{}, fmt.Errorf("error occurred")
}
