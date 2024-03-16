package api

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/image-storage-service/cmd/api"
	"github.com/shared/mocks"
	"github.com/shared/models"
	"github.com/stretchr/testify/require"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUploadFileShouldReturn200(t *testing.T) {
	server := initServer(&mocks.MockImageStorageLogic{})

	body, contentType, err := getRequestBody()
	if err != nil {
		t.Fatalf("Error creating requestBody: %s", err)
	}

	response := sendRequestUpload(t, server, body, contentType)

	require.Equal(t, http.StatusOK, response.Code)
}

func TestUploadFileShouldReturn500WhenLogicReturnsError(t *testing.T) {
	server := initServer(&mocks.MockImageStorageLogicReturnsError{})

	body, contentType, err := getRequestBody()
	if err != nil {
		t.Fatalf("Error creating requestBody: %s", err)
	}

	response := sendRequestUpload(t, server, body, contentType)

	require.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestGetIsoFileShouldReturn200(t *testing.T) {
	server := initServer(&mocks.MockImageStorageLogic{})

	response := sendRequestDownload(t, server, "test.txt")

	require.Equal(t, http.StatusOK, response.Code)
}

func TestGetIsoFileShouldReturn404WhenFileDoesNotExist(t *testing.T) {
	server := initServer(&mocks.MockImageStorageLogic{})

	response := sendRequestDownload(t, server, "test2.txt")

	require.Equal(t, http.StatusNotFound, response.Code)
}

func TestGetIsoFileShouldReturn500WhenLogicReturnsError(t *testing.T) {
	server := initServer(&mocks.MockImageStorageLogicReturnsError{})

	response := sendRequestDownload(t, server, "test.txt")

	require.Equal(t, http.StatusInternalServerError, response.Code)
}

func getRequestBody() (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileField, err := writer.CreateFormFile("file", "qwer123.iso")
	if err != nil {
		return nil, "", fmt.Errorf("error creating form file: %s", err)
	}

	_, err = fileField.Write([]byte("example ISO image data"))
	if err != nil {
		return nil, "", fmt.Errorf("error writing mock ISO image data: %s", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, "", fmt.Errorf("error closing form: %s", err)
	}

	return body, writer.FormDataContentType(), nil
}

func initServer(imageStorageLogic models.ImageStorageLogicInterface) *api.GinServer {
	server := &api.GinServer{
		ImageStorageLogic: imageStorageLogic,
	}
	return server
}

func sendRequestUpload(t *testing.T, server *api.GinServer, payload *bytes.Buffer, contentType string) *httptest.ResponseRecorder {
	route := gin.Default()
	route.POST("/upload", server.UploadFile)

	request, err := http.NewRequest("POST", "/upload", payload)
	if err != nil {
		t.Errorf("error while sending request")
	}

	request.Header.Set("Content-Type", contentType)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)

	return recorder
}

func sendRequestDownload(t *testing.T, server *api.GinServer, fileName string) *httptest.ResponseRecorder {
	route := gin.Default()
	route.GET("/download/:fileName", func(c *gin.Context) {
		server.GetIsoFile(c, c.Param("fileName"))
	})

	payloadByte := []byte("")

	request, err := http.NewRequest("GET", "/download/"+fileName, bytes.NewBuffer(payloadByte))
	if err != nil {
		t.Errorf("error while sending request")
	}

	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)

	return recorder
}
