package mocks

import (
	"fmt"
	"github.com/shared/models"
)

type MockImageBuilderLogic struct{}

func (m *MockImageBuilderLogic) BuildImage(_ models.ImageConfig) (models.ImageId, error) {
	return "WZ3h633-p", nil
}

type MockImageBuilderLogicReturnsError struct{}

func (m *MockImageBuilderLogicReturnsError) BuildImage(_ models.ImageConfig) (models.ImageId, error) {
	return "", fmt.Errorf("error occurred")
}
