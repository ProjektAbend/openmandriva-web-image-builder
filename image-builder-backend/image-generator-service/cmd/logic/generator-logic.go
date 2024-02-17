package logic

import (
	"encoding/json"
	"fmt"
	"github.com/shared/constants"
	"github.com/shared/models"
	"log"
	"os/exec"
	"time"
)

type GeneratorLogic struct {
	MessageBroker models.MessageBrokerInterface
}

func (c *GeneratorLogic) ProcessBuildRequests() {
	for {
		imageConfig, isEmpty, err := c.ProcessBuildRequest()
		if err != nil {
			log.Printf("error while consuming message: %s", err)
		}
		if !isEmpty {
			generateImage(imageConfig)
		}
	}
}

func (c *GeneratorLogic) ProcessBuildRequest() (models.ImageConfig, bool, error) {
	message, err := c.MessageBroker.ConsumeMessage(constants.BUILD_QUEUE)
	if err != nil {
		return models.ImageConfig{}, false, fmt.Errorf("error consuming message: %s", err)
	}

	if message.Body == nil {
		return models.ImageConfig{}, true, nil
	}

	var imageConfig models.ImageConfig
	err = json.Unmarshal(message.Body, &imageConfig)
	if err != nil {
		return models.ImageConfig{}, false, fmt.Errorf("error unmarshalling message: %s", err)
	}

	return imageConfig, false, nil
}

func generateImage(imageConfig models.ImageConfig) {
	log.Printf("Processing image with ID: %v", *imageConfig.ImageId)
	output, err := runCommand("./os-image-builder/build", "-h")
	if err != nil {
		log.Printf("error running command: %s", err)
	}
	log.Printf("output: %s", output)
	time.Sleep(5 * time.Second)
}

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	outputStr := string(output)
	return outputStr, nil
}
