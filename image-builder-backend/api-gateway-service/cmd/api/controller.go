package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net"
	"net/http"
)

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type ImageBuilder interface {
	BuildImage(imageConfig ImageConfig) (ImageInfo, error)
}

type GinServer struct {
	ImageBuilder ImageBuilder
}

func (s GinServer) BuildImage(context *gin.Context) {
	var imageConfig ImageConfig
	if err := context.BindJSON(&imageConfig); err != nil {
		return
	}

	validate := validator.New()
	if err := validate.Struct(imageConfig); err != nil {
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	imageInfo, err := s.ImageBuilder.BuildImage(imageConfig)
	if err != nil {
		log.Printf("Error in BuildImage: %s", err)
		sendError(context, http.StatusInternalServerError, "Failed to build the image")
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"imageId":        imageInfo.ImageId,
		"isAvailable":    imageInfo.IsAvailable,
		"availableUntil": imageInfo.AvailableUntil,
	})
}

func (s GinServer) GetImageById(c *gin.Context, imageId ImageId) {
	// TODO: implement
}

func (s GinServer) GetStatusOfImageById(c *gin.Context, imageId ImageId) {
	// TODO: implement
}

func sendError(c *gin.Context, code int, message string) {
	err := Error{
		Code:    int32(code),
		Message: message,
	}
	c.JSON(code, err)
}

func StartServer(imageBuilder ImageBuilder) {
	route := gin.Default()
	server := &GinServer{ImageBuilder: imageBuilder}
	RegisterHandlers(route, server)

	s := &http.Server{
		Handler: route,
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
	}
	s.ListenAndServe()
}
