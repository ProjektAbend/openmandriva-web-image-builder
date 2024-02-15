package logic_test

import (
	"github.com/api-gateway-service/cmd/api"
	"github.com/api-gateway-service/cmd/logic"
	"github.com/api-gateway-service/test/mocks"
	"github.com/stretchr/testify/require"
	"github.com/teris-io/shortid"
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
	shortIdGenerator, _ := shortid.New(1, shortid.DefaultABC, 2342)

	imageBuilderLogic := &logic.ImageBuilderLogic{
		MessageBroker:    messageBroker,
		ShortIdGenerator: shortIdGenerator,
	}

	return imageBuilderLogic
}
