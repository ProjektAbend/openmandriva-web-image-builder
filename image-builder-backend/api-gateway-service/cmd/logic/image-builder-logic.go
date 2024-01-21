package logic

import (
	"fmt"
	"github.com/api-gateway-service/cmd/api"
)

type ImageBuilderLogic struct{}

func (c *ImageBuilderLogic) BuildImage(imageConfig api.ImageConfig) error {
	// TODO: send build request to message broker
	fmt.Println(imageConfig)
	return nil
}
