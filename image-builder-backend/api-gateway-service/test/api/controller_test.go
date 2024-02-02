package api

import (
	"bytes"
	"github.com/api-gateway-service/cmd/api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockImageBuilder struct{}

var expectedImageId = "a1b2c3"

func (m *mockImageBuilder) BuildImage(imageConfig api.ImageConfig) (api.ImageId, error) {
	return expectedImageId, nil
}

func sendRequestBuild(payload string) (*httptest.ResponseRecorder, error) {
	validate := validator.New()
	server := &api.GinServer{
		ImageBuilder: &mockImageBuilder{},
		Validate:     validate,
	}
	route := gin.Default()
	route.POST("/whatever", server.BuildImage)

	payloadByte := []byte(payload)

	request, err := http.NewRequest("POST", "/whatever", bytes.NewBuffer(payloadByte))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)

	return recorder, nil
}

func TestBuildImageShouldReturn201WhenCorrectImageConfig(t *testing.T) {
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

	response, err := sendRequestBuild(requestBody)
	if err != nil {
		t.Errorf("error while sending request")
	}

	if status := response.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestBuildImageShouldAlsoReturnImageIdWhenReturning201(t *testing.T) {
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

	response, err := sendRequestBuild(requestBody)
	if err != nil {
		t.Errorf("error while sending request")
	}

	if status := response.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expectedResponse := "{\"imageId\": \"" + expectedImageId + "\"}"
	actualResponse := response.Body.String()

	require.JSONEq(t, expectedResponse, actualResponse)
}

func TestBuildImageShouldReturn201WhenImageConfigHasOnlyArchitecture(t *testing.T) {
	requestBody := `{
        "architecture": "aarch64-uefi"
    }`

	response, err := sendRequestBuild(requestBody)
	if err != nil {
		t.Errorf("error while sending request")
	}

	if status := response.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestBuildImageShouldReturn400WhenArchitectureOfImageConfigIsEmpty(t *testing.T) {
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

	response, err := sendRequestBuild(requestBody)
	if err != nil {
		t.Errorf("error while sending request")
	}

	if status := response.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestBuildImageShouldReturn400WhenArchitectureOfImageConfigIsMissing(t *testing.T) {
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

	response, err := sendRequestBuild(requestBody)
	if err != nil {
		t.Errorf("error while sending request")
	}

	if status := response.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
