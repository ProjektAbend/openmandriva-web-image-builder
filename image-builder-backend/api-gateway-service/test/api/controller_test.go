package api

import (
	"bytes"
	"encoding/json"
	"github.com/api-gateway-service/cmd/api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockImageBuilder struct{}

var expectedImageInfo = api.ImageInfo{
	ImageId: "a1b2c3",
}

func (m *mockImageBuilder) BuildImage(imageConfig api.ImageConfig) (api.ImageInfo, error) {
	return expectedImageInfo, nil
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
		return &httptest.ResponseRecorder{}, err
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

func TestBuildImageShouldAlsoReturnValidImageInfoWhenReturning201(t *testing.T) {
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

	var responseImageInfo api.ImageInfo
	if err := json.Unmarshal(response.Body.Bytes(), &responseImageInfo); err != nil {
		t.Errorf("Failed to parse response JSON: %v", err)
		return
	}

	expectedImageInfo2 := `{
		"availableUntil": null,
		"imageId":        "a1b2c3",
		"isAvailable":    null
	}`

	actualImageInfo := response.Body.String()

	var data1, data2 interface{}

	if err := json.Unmarshal([]byte(expectedImageInfo2), &data1); err != nil {
		t.Errorf("Error parsing expectedImageInfo: %s", err)
		return
	}

	if err := json.Unmarshal([]byte(actualImageInfo), &data2); err != nil {
		t.Errorf("Error parsing actualImageInfo: %s", err)
		return
	}

	json1, _ := json.Marshal(data1)
	json2, _ := json.Marshal(data2)

	if string(json1) != string(json2) {
		t.Errorf("Response expected to be %s but was %s", json1, json2)
	}
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
