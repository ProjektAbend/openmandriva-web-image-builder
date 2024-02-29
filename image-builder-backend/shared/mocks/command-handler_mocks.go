package mocks

import (
	"fmt"
	"github.com/shared/models"
	"log"
	"os"
	"path/filepath"
	"time"
)

type MockCommandHandler struct{}

func (c *MockCommandHandler) GenerateImage(_ models.ImageConfig) error {
	return nil
}

type MockCommandHandlerGenerateFakeIso struct{}

func (c *MockCommandHandlerGenerateFakeIso) GenerateImage(imageConfig models.ImageConfig) error {
	log.Printf("Processing image with ID: %v", *imageConfig.ImageId)
	time.Sleep(5 * time.Second)
	err := createFakeIso(*imageConfig.ImageId)
	if err != nil {
		return fmt.Errorf("error while creating fake iso: %s", err)
	}
	return nil
}

func createFakeIso(fileName string) error {
	outputDir := "generated-images"
	filePath := filepath.Join(outputDir, fileName)
	file, err := os.Create(filePath + ".iso")
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
