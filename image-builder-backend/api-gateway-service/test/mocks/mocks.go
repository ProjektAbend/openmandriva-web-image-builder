package mocks

import (
	"errors"
	"github.com/api-gateway-service/cmd/api"
)

type ImageBuilderLogic interface {
	BuildImage(imageConfig api.ImageConfig) (api.ImageId, error)
}

type MockImageBuilderLogic struct{}

func (m *MockImageBuilderLogic) BuildImage(imageConfig api.ImageConfig) (api.ImageId, error) {
	return "a1b2c3", nil
}

type MockImageBuilderLogicReturnsError struct{}

func (m *MockImageBuilderLogicReturnsError) BuildImage(imageConfig api.ImageConfig) (api.ImageId, error) {
	return "", errors.New("error occurred")
}

type MessageBroker interface {
	SendMessageToQueue(message string, queue string) error
}

type MockMessageBroker struct{}

func (m *MockMessageBroker) SendMessageToQueue(message string, queue string) error {
	return nil
}

type MockMessageBrokerReturnsError struct{}

func (m *MockMessageBrokerReturnsError) SendMessageToQueue(message string, queue string) error {
	return errors.New("error occurred")
}
