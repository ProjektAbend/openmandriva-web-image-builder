package mocks

import (
	"errors"
	"github.com/api-gateway-service/cmd/api"
)

type ImageBuilder interface {
	BuildImage(imageConfig api.ImageConfig) (api.ImageId, error)
}

type MockImageBuilder struct{}

func (m *MockImageBuilder) BuildImage(imageConfig api.ImageConfig) (api.ImageId, error) {
	return "a1b2c3", nil
}

type MockImageBuilderReturnsError struct{}

func (m *MockImageBuilderReturnsError) BuildImage(imageConfig api.ImageConfig) (api.ImageId, error) {
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
