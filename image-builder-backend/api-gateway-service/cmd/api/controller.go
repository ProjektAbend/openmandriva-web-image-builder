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

type GinServer struct{}

func (s GinServer) BuildImage(c *gin.Context) {
	// TODO: implement
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

func StartServer() {
	r := gin.Default()
	RegisterHandlers(r, GinServer{})
	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
	}
	s.ListenAndServe()
}
