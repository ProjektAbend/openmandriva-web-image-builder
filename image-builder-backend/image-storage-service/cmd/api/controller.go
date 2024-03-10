package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shared/models"
	"log"
	"net"
	"net/http"
)

type GinServer struct {
	ImageStorageLogic models.ImageStorageLogicInterface
}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (s GinServer) UploadFile(context *gin.Context) {
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		context.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	defer file.Close()

	err = s.ImageStorageLogic.StoreImage(file, header.Filename)
	if err != nil {
		log.Printf("Error in StoreImage: %s", err)
		sendError(context, http.StatusInternalServerError, "Failed to save the file.")
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "file was successfully saved to disk."})
}

func (s GinServer) GetIsoFile(context *gin.Context, fileName string) {
	filePath, err := s.ImageStorageLogic.GetIsoFile(fileName)
	if err != nil {
		log.Printf("Error in GetIsoFile: %s", err)
		sendError(context, http.StatusInternalServerError, "Failed to retrieve the file")
		return
	}

	if !s.ImageStorageLogic.DoesFileExist(filePath) {
		log.Printf("file %s not found.", fileName)
		sendError(context, http.StatusNotFound, "file not found.")
		return
	}
	context.File(filePath)
}

func sendError(c *gin.Context, code int, message string) {
	err := Error{
		Code:    int32(code),
		Message: message,
	}
	c.JSON(code, err)
}

func StartServer(imageStorageLogic models.ImageStorageLogicInterface) {
	route := gin.Default()
	server := &GinServer{
		ImageStorageLogic: imageStorageLogic,
	}
	RegisterHandlers(route, server)

	s := &http.Server{
		Handler: route,
		Addr:    net.JoinHostPort("0.0.0.0", "8084"),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Printf("Error starting Server: %s", err)
	}
}
