package api_test

import (
	"bytes"
	"github.com/api-gateway-service/cmd/api"
	"github.com/api-gateway-service/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuildImageShouldReturn201WhenCorrectImageConfig(t *testing.T) {
	server := initServer(&mocks.MockImageBuilderLogic{})

	requestBody := `{
        "architecture": "aarch64-uefi",
        "version": "4.2",
        "desktop": "kde",
        "services": [
            {
                "name": "cloud-init",
                "disabled": true
            }
        ],
        "packages": [
            {
                "name": "vim-enhanced",
                "installWeakDependencies": true,
                "packageType": "INDIVIDUAL",
                "repositoryType": "UNSUPPORTED"
            }
        ]
    }`

	response := sendRequestBuild(t, server, requestBody)

	require.Equal(t, http.StatusCreated, response.Code)
}

func TestBuildImageShouldAlsoReturnImageIdWhenReturning201(t *testing.T) {
	server := initServer(&mocks.MockImageBuilderLogic{})

	requestBody := `{
        "architecture": "aarch64-uefi",
        "version": "4.2",
        "desktop": "kde",
        "services": [
            {
                "name": "cloud-init",
                "disabled": true
            }
        ],
        "packages": [
            {
                "name": "vim-enhanced",
                "installWeakDependencies": true,
                "packageType": "INDIVIDUAL",
                "repositoryType": "UNSUPPORTED"
            }
        ]
    }`

	response := sendRequestBuild(t, server, requestBody)

	require.Equal(t, http.StatusCreated, response.Code)

	expectedResponse := "{\"imageId\": \"WZ3h633-p\"}"
	actualResponse := response.Body.String()

	require.JSONEq(t, expectedResponse, actualResponse)
}

func TestBuildImageShouldReturn201WhenImageConfigHasOnlyArchitecture(t *testing.T) {
	server := initServer(&mocks.MockImageBuilderLogic{})

	requestBody := `{
        "architecture": "aarch64-uefi"
    }`

	response := sendRequestBuild(t, server, requestBody)

	require.Equal(t, http.StatusCreated, response.Code)
}

func TestBuildImageShouldReturn400WhenArchitectureOfImageConfigIsEmpty(t *testing.T) {
	server := initServer(&mocks.MockImageBuilderLogic{})

	requestBody := `{
        "architecture": "",
        "version": "4.2",
        "desktop": "kde",
        "services": [
            {
                "name": "cloud-init",
                "disabled": true
            }
        ],
        "packages": [
            {
                "name": "vim-enhanced",
                "installWeakDependencies": true,
                "packageType": "INDIVIDUAL",
                "repositoryType": "UNSUPPORTED"
            }
        ]
    }`

	response := sendRequestBuild(t, server, requestBody)

	require.Equal(t, http.StatusBadRequest, response.Code)
}

func TestBuildImageShouldReturn400WhenRequestBodyIsEmpty(t *testing.T) {
	server := initServer(&mocks.MockImageBuilderLogic{})

	requestBody := ""

	response := sendRequestBuild(t, server, requestBody)

	require.Equal(t, http.StatusBadRequest, response.Code)
}

func TestBuildImageShouldReturn400WhenArchitectureOfImageConfigIsMissing(t *testing.T) {
	server := initServer(&mocks.MockImageBuilderLogic{})

	requestBody := `{
        "version": "4.2",
        "desktop": "kde",
        "services": [
            {
                "name": "cloud-init",
                "disabled": true
            }
        ],
        "packages": [
            {
                "name": "vim-enhanced",
                "installWeakDependencies": true,
                "packageType": "INDIVIDUAL",
                "repositoryType": "UNSUPPORTED"
            }
        ]
    }`

	response := sendRequestBuild(t, server, requestBody)

	require.Equal(t, http.StatusBadRequest, response.Code)
}

func TestBuildImageShouldReturn500WhenLogicReturnsError(t *testing.T) {
	server := initServer(&mocks.MockImageBuilderLogicReturnsError{})

	requestBody := `{
        "architecture": "aarch64-uefi",
        "version": "4.2",
        "desktop": "kde",
        "services": [
            {
                "name": "cloud-init",
                "disabled": true
            }
        ],
        "packages": [
            {
                "name": "vim-enhanced",
                "installWeakDependencies": true,
                "packageType": "INDIVIDUAL",
                "repositoryType": "UNSUPPORTED"
            }
        ]
    }`

	response := sendRequestBuild(t, server, requestBody)

	require.Equal(t, http.StatusInternalServerError, response.Code)
}

func initServer(imageBuilder api.ImageBuilderLogicInterface) *api.GinServer {
	validate := validator.New()
	server := &api.GinServer{
		ImageBuilderLogic: imageBuilder,
		Validate:          validate,
	}
	return server
}

func sendRequestBuild(t *testing.T, server *api.GinServer, payload string) *httptest.ResponseRecorder {
	route := gin.Default()
	route.POST("/whatever", server.BuildImage)

	payloadByte := []byte(payload)

	request, err := http.NewRequest("POST", "/whatever", bytes.NewBuffer(payloadByte))
	if err != nil {
		t.Errorf("error while sending request")
	}

	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)

	return recorder
}
