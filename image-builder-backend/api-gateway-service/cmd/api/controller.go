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
	BuildImage(imageConfig ImageConfig) (ImageId, error)
}

type GinServer struct {
	ImageBuilder ImageBuilder
	Validate     *validator.Validate
}

func (s GinServer) BuildImage(context *gin.Context) {
	var imageConfig ImageConfig
	if err := context.BindJSON(&imageConfig); err != nil {
		return
	}

	if err := s.Validate.Struct(imageConfig); err != nil {
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	imageId, err := s.ImageBuilder.BuildImage(imageConfig)
	if err != nil {
		log.Printf("Error in BuildImage: %s", err)
		sendError(context, http.StatusInternalServerError, "Failed to build the image")
		return
	}
	context.JSON(http.StatusCreated, gin.H{"imageId": imageId})
}

func (s GinServer) GetImageById(context *gin.Context, imageId ImageId) {
	context.Status(http.StatusNotImplemented)
}

func (s GinServer) GetStatusOfImageById(context *gin.Context, imageId ImageId) {
	context.Status(http.StatusNotImplemented)
}

func sendError(c *gin.Context, code int, message string) {
	err := Error{
		Code:    int32(code),
		Message: message,
	}
	c.JSON(code, err)
}

func StartServer(imageBuilder ImageBuilder, validate *validator.Validate) {
	route := gin.Default()
	server := &GinServer{
		ImageBuilder: imageBuilder,
		Validate:     validate,
	}
	RegisterHandlers(route, server)

	s := &http.Server{
		Handler: route,
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
	}
	s.ListenAndServe()
}
