package main

import (
	"flag"
	"github.com/ProjektAbend/api-gateway-service/api"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type GinServer struct{}

func (s GinServer) BuildImage(c *gin.Context) {
	// TODO: implement
}

func (s GinServer) GetImageById(c *gin.Context, imageId api.ImageId) {
	// TODO: implement
}

func (s GinServer) GetStatusOfImageById(c *gin.Context, imageId api.ImageId) {
	// TODO: implement
}

func main() {
	r := gin.Default()
	g := GinServer{}
	api.RegisterHandlers(r, g)
	port := flag.String("port", "8080", "Port for test HTTP server")
	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", *port),
	}
	s.ListenAndServe()
}
