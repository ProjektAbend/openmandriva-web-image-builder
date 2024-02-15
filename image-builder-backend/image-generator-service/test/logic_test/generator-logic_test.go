package logic_test

import (
	"github.com/image-generator-service/cmd/api"
	"github.com/image-generator-service/cmd/logic"
	"github.com/image-generator-service/test/mocks"
	messagebroker "github.com/shared/message-broker"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProcessBuildRequestShouldReturnCorrectImageConfig(t *testing.T) {
	generatorLogic := initGeneratorLogic(&mocks.MockMessageBroker{})

	expectedImageConfig := api.ImageConfig{
		Architecture: "aarch64-uefi",
	}

	actualImageConfig, _ := generatorLogic.ProcessBuildRequest()

	require.Equal(t, expectedImageConfig, actualImageConfig)
}

func TestProcessBuildRequestShouldErrorWhenQueueIsEmpty(t *testing.T) {
	generatorLogic := initGeneratorLogic(&mocks.MockMessageBrokerReturnsError{})

	_, err := generatorLogic.ProcessBuildRequest()

	require.NotEqual(t, nil, err)
}

func initGeneratorLogic(messageBroker messagebroker.MessageBrokerInterface) *logic.GeneratorLogic {
	generatorLogic := &logic.GeneratorLogic{
		MessageBroker: messageBroker,
	}
	return generatorLogic
}
