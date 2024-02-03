package logic

import (
	"github.com/api-gateway-service/cmd/api"
	"github.com/api-gateway-service/cmd/logic"
	"github.com/api-gateway-service/test/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuildImageShouldReturnString(t *testing.T) {
	imageBuilderLogic := initImageBuilderLogic(&mocks.MockMessageBroker{})

	imageId, _ := imageBuilderLogic.BuildImage(api.ImageConfig{})

	require.IsType(t, "", imageId)
}

func TestBuildImageShouldReturnErrorWhenMessageBrokerFailed(t *testing.T) {
	imageBuilderLogic := initImageBuilderLogic(&mocks.MockMessageBrokerReturnsError{})

	_, err := imageBuilderLogic.BuildImage(api.ImageConfig{})

	require.NotEqual(t, nil, err)
}

func initImageBuilderLogic(messageBroker mocks.MessageBroker) *logic.ImageBuilderLogic {
	imageBuilderLogic := &logic.ImageBuilderLogic{
		MessageBroker: messageBroker,
	}
	return imageBuilderLogic
}
