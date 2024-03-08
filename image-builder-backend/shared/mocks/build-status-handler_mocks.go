package mocks

import (
	"fmt"
	"github.com/shared/models"
)

type MockBuildStatusHandler struct{}

func (_ *MockBuildStatusHandler) SetStatusOfImageBuild(_ models.ImageId, _ models.Status) {
	fmt.Printf("mock func")
}
