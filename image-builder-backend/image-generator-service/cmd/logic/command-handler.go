package logic

import (
	"fmt"
	"github.com/shared/models"
	"log"
	"time"
)

type CommandHandler struct{}

func (c *CommandHandler) GenerateImage(imageConfig models.ImageConfig) error {
	log.Printf("Processing image with ID: %v", *imageConfig.ImageId)
	err := c.runCommand("./os-image-builder/build", imageConfig.Architecture)
	if err != nil {
		return fmt.Errorf("error running command %s", err)
	}
	return nil
}

func (c *CommandHandler) runCommand(_ string, _ ...string) error {
	// TODO: implement
	time.Sleep(5 * time.Second)
	return nil
}
