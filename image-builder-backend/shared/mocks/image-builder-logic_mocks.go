package mocks

import (
	"fmt"
	"github.com/shared/models"
)

type MockImageBuilderLogic struct{}

func (_ *MockImageBuilderLogic) BuildImage(_ models.ImageConfig) (models.ImageId, error) {
	return "WZ3h633-p", nil
}

func (_ *MockImageBuilderLogic) GetStatusOfImage(imageId models.ImageId) (models.ImageInfo, error) {
	status := models.ACCEPTED
	return models.ImageInfo{
		AvailableUntil: nil,
		ImageId:        imageId,
		Status:         &status,
	}, nil
}

type MockImageBuilderLogicReturnsError struct{}

func (_ *MockImageBuilderLogicReturnsError) BuildImage(_ models.ImageConfig) (models.ImageId, error) {
	return "", fmt.Errorf("error occurred")
}

func (_ *MockImageBuilderLogicReturnsError) GetStatusOfImage(_ models.ImageId) (models.ImageInfo, error) {
	return models.ImageInfo{}, fmt.Errorf("error occurred")
}
