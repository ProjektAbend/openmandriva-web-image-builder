package api

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type ImageBuilder interface {
	BuildImage(imageConfig ImageConfig) error
	// Add other methods as needed
}

type GinServer struct {
	ImageBuilder ImageBuilder
}

func (s GinServer) BuildImage(c *gin.Context) {
	// map json from request body to struct
	var imageConfig ImageConfig
	if err := c.BindJSON(&imageConfig); err != nil {
		return
	}

	err := s.ImageBuilder.BuildImage(imageConfig)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Failed to build the image")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "The build process for the desired image has started."})
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
	r := gin.Default()
	server := &GinServer{ImageBuilder: imageBuilder}
	RegisterHandlers(r, server)

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
	}
	s.ListenAndServe()
}
