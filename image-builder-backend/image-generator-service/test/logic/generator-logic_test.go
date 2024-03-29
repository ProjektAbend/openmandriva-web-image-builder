package logic

import (
	"github.com/image-generator-service/cmd/logic"
	"github.com/shared/mocks"
	"github.com/shared/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProcessBuildRequestShouldReturnCorrectImageConfig(t *testing.T) {
	generatorLogic := initGeneratorLogic(&mocks.MockMessageBroker{})

	imageId := "WZ3h633-p"
	expectedImageConfig := models.ImageConfig{
		Architecture: "aarch64-uefi",
		ImageId:      &imageId,
	}

	actualImageConfig, isEmpty, _ := generatorLogic.GetImageConfig()

	require.False(t, isEmpty)
	require.Equal(t, expectedImageConfig, actualImageConfig)
}

func TestProcessBuildRequestShouldErrorWhenMessageBrokerErrors(t *testing.T) {
	generatorLogic := initGeneratorLogic(&mocks.MockMessageBrokerReturnsError{})

	_, isEmpty, err := generatorLogic.GetImageConfig()

	require.NotEqual(t, nil, err)
	require.False(t, isEmpty)
}

func TestProcessBuildRequestShouldReturnTrueWhenQueueIsEmpty(t *testing.T) {
	generatorLogic := initGeneratorLogic(&mocks.MockMessageBrokerHasEmptyQueue{})

	_, isEmpty, err := generatorLogic.GetImageConfig()

	require.Equal(t, nil, err)
	require.True(t, isEmpty)
}

func initGeneratorLogic(messageBroker models.MessageBrokerInterface) *logic.GeneratorLogic {
	commandHandler := &mocks.MockCommandHandler{}
	buildStatusHandler := &mocks.MockBuildStatusHandler{}
	generatorLogic := &logic.GeneratorLogic{
		MessageBroker:      messageBroker,
		CommandHandler:     commandHandler,
		BuildStatusHandler: buildStatusHandler,
	}
	return generatorLogic
}
