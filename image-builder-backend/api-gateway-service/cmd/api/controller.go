package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shared/models"
	"log"
	"net"
	"net/http"
)

type GinServer struct {
	ImageBuilderLogic models.ImageBuilderLogicInterface
	Validate          *validator.Validate
}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (s GinServer) BuildImage(context *gin.Context) {
	var imageConfig models.ImageConfig
	if err := context.BindJSON(&imageConfig); err != nil {
		return
	}

	if err := s.Validate.Struct(imageConfig); err != nil {
		log.Printf("Error validating imageConfig: %s", err)
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	imageId, err := s.ImageBuilderLogic.BuildImage(imageConfig)
	if err != nil {
		log.Printf("Error in BuildImage: %s", err)
		sendError(context, http.StatusInternalServerError, "Failed to build the image")
		return
	}
	context.JSON(http.StatusCreated, gin.H{"imageId": imageId})
}

func (s GinServer) GetImageById(context *gin.Context, imageId models.ImageId) {
	fileContent, err := s.ImageBuilderLogic.GetImage(imageId)
	if err != nil {
		log.Printf("Failed to download file: %s", err)
		sendError(context, http.StatusInternalServerError, "Failed to retrieve the image")
		return
	}

	context.Header("Content-Type", "application/x-iso9660-image")
	context.Data(http.StatusOK, "application/x-iso9660-image", fileContent)
}

func (s GinServer) GetStatusOfImageById(context *gin.Context, imageId models.ImageId) {
	imageInfo, err := s.ImageBuilderLogic.GetStatusOfImage(imageId)
	if err != nil {
		log.Printf("Failed to retrieve status of %s: %s", imageId, err)
		sendError(context, http.StatusInternalServerError, "Failed to retrieve status")
		return
	}
	context.JSON(http.StatusOK, imageInfo)
}

func sendError(c *gin.Context, code int, message string) {
	err := Error{
		Code:    int32(code),
		Message: message,
	}
	c.JSON(code, err)
}

func StartServer(imageBuilder models.ImageBuilderLogicInterface, validate *validator.Validate) {
	route := gin.Default()
	server := &GinServer{
		ImageBuilderLogic: imageBuilder,
		Validate:          validate,
	}
	RegisterHandlers(route, server)

	s := &http.Server{
		Handler: route,
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Printf("Error starting Server: %s", err)
	}
}
