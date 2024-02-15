package mocks

import (
	"errors"
	"github.com/api-gateway-service/cmd/api"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MockImageBuilderLogic struct{}

func (m *MockImageBuilderLogic) BuildImage(_ api.ImageConfig) (api.ImageId, error) {
	return "WZ3h633-p", nil
}

type MockImageBuilderLogicReturnsError struct{}

func (m *MockImageBuilderLogicReturnsError) BuildImage(_ api.ImageConfig) (api.ImageId, error) {
	return "", errors.New("error occurred")
}

type MockMessageBroker struct{}

func (_ *MockMessageBroker) SendMessageToQueue(_ string, _ string) error {
	return nil
}

func (_ *MockMessageBroker) ConsumeMessage(_ string) (amqp.Delivery, error) {
	return amqp.Delivery{}, nil
}

type MockMessageBrokerReturnsError struct{}

func (_ *MockMessageBrokerReturnsError) SendMessageToQueue(_ string, _ string) error {
	return errors.New("error occurred")
}

func (_ *MockMessageBrokerReturnsError) ConsumeMessage(_ string) (amqp.Delivery, error) {
	return amqp.Delivery{}, nil
}
