package main

import (
	"flag"
	"github.com/ProjektAbend/api-gateway-service/api"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

func NewGinServer(serverInterface *api.ServerInterface, port string) *http.Server {
	r := gin.Default()

	api.RegisterHandlers(r, serverInterface)

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}
	return s
}

func main() {
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()
	// Create an instance of our handler which satisfies the generated interface
	serverInterface := api.ServerInterface()
	s := NewGinServer(serverInterface, *port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
