package logic_test

import (
	"github.com/api-gateway-service/cmd/logic"
	"github.com/shared/mocks"
	"github.com/shared/models"
	"github.com/stretchr/testify/require"
	"github.com/teris-io/shortid"
	"testing"
)

func TestBuildImageShouldReturnString(t *testing.T) {
	imageBuilderLogic := initImageBuilderLogic(&mocks.MockMessageBroker{})

	imageId, _ := imageBuilderLogic.BuildImage(models.ImageConfig{})

	require.IsType(t, "", imageId)
}

func TestBuildImageShouldReturnErrorWhenMessageBrokerFailed(t *testing.T) {
	imageBuilderLogic := initImageBuilderLogic(&mocks.MockMessageBrokerReturnsError{})

	_, err := imageBuilderLogic.BuildImage(models.ImageConfig{})

	require.NotEqual(t, nil, err)
}

func initImageBuilderLogic(messageBroker models.MessageBrokerInterface) *logic.ImageBuilderLogic {
	shortIdGenerator, _ := shortid.New(1, shortid.DefaultABC, 2342)

	buildStatusHandler := &mocks.MockBuildStatusHandler{}

	imageBuilderLogic := &logic.ImageBuilderLogic{
		MessageBroker:      messageBroker,
		ShortIdGenerator:   shortIdGenerator,
		BuildStatusHandler: buildStatusHandler,
	}

	return imageBuilderLogic
}
