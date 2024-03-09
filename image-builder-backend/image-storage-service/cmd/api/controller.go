package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

type GinServer struct{}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (s GinServer) UploadFile(context *gin.Context) {
	context.Status(http.StatusNotImplemented)
}

func StartServer() {
	route := gin.Default()
	server := &GinServer{}
	RegisterHandlers(route, server)

	s := &http.Server{
		Handler: route,
		Addr:    net.JoinHostPort("0.0.0.0", "8084"),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Printf("Error starting Server: %s", err)
	}
}
