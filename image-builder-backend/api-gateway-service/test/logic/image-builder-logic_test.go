package logic

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

func TestBuildImageShouldReturnErrorWhenMessageBrokerFails(t *testing.T) {
	imageBuilderLogic := initImageBuilderLogic(&mocks.MockMessageBrokerReturnsError{})

	_, err := imageBuilderLogic.BuildImage(models.ImageConfig{})

	require.NotEqual(t, nil, err)
}

func TestGetStatusOfImageShouldReturnCorrectImageInfo(t *testing.T) {
	imageBuilderLogic := initImageBuilderLogic(&mocks.MockMessageBroker{})

	imageId := "WZ3h633-p"
	status := models.ACCEPTED

	expectedImageInfo := models.ImageInfo{
		AvailableUntil: nil,
		ImageId:        imageId,
		Status:         &status,
	}

	actualImageInfo, _ := imageBuilderLogic.GetStatusOfImage(imageId)

	require.Equal(t, expectedImageInfo, actualImageInfo)
}

func TestGetStatusOfImageShouldReturnDoesNotExistStatus(t *testing.T) {
	imageBuilderLogic := initImageBuilderLogic(&mocks.MockMessageBrokerHasEmptyQueue{})

	imageId := "WZ3h633-p"
	status := models.DOESNOTEXIST

	expectedImageInfo := models.ImageInfo{
		AvailableUntil: nil,
		ImageId:        imageId,
		Status:         &status,
	}

	actualImageInfo, _ := imageBuilderLogic.GetStatusOfImage(imageId)

	require.Equal(t, expectedImageInfo, actualImageInfo)
}

func TestGetStatusOfImageShouldReturnErrorWhenMessageBrokerFails(t *testing.T) {
	imageBuilderLogic := initImageBuilderLogic(&mocks.MockMessageBrokerReturnsError{})

	imageId := "WZ3h633-p"

	_, err := imageBuilderLogic.GetStatusOfImage(imageId)

	require.NotEqual(t, nil, err)
}

func TestFindLatestStatusShouldReturnAvailable(t *testing.T) {
	expected := models.AVAILABLE

	var statuses []models.Status

	statuses = append(statuses, models.REQUESTED)
	statuses = append(statuses, models.ACCEPTED)
	statuses = append(statuses, models.AVAILABLE)
	statuses = append(statuses, models.UPLOADSTARTED)

	actual := logic.FindLatestStatus(statuses)

	require.Equal(t, expected, actual)
}

func TestFindLatestStatusShouldReturnUploadStarted(t *testing.T) {
	expected := models.UPLOADSTARTED

	var statuses []models.Status

	statuses = append(statuses, models.REQUESTED)
	statuses = append(statuses, models.UPLOADSTARTED)
	statuses = append(statuses, models.ACCEPTED)
	statuses = append(statuses, models.FETCHED)

	actual := logic.FindLatestStatus(statuses)

	require.Equal(t, expected, actual)
}

func TestFindLatestStatusShouldReturnDoesNotExistStatusWhenSliceIsEmpty(t *testing.T) {
	expected := models.DOESNOTEXIST

	actual := logic.FindLatestStatus([]models.Status{})

	require.Equal(t, expected, actual)
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
