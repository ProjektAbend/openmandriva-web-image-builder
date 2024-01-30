package api

import (
	"bytes"
	"github.com/api-gateway-service/cmd/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockImageBuilder struct{}

func (m *mockImageBuilder) BuildImage(imageConfig api.ImageConfig) error {
	// Simulate the behavior of BuildImage method
	return nil
}

func TestBuildImageShouldReturn201WhenCorrectImageConfig(t *testing.T) {
	server := &api.GinServer{ImageBuilder: &mockImageBuilder{}}
	route := gin.Default()
	route.POST("/build", server.BuildImage)

	payload := []byte(`{
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
    }`)

	request, err := http.NewRequest("POST", "/build", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestBuildImageShouldReturn400WhenArchitectureOfImageConfigIsEmpty(t *testing.T) {
	server := &api.GinServer{ImageBuilder: &mockImageBuilder{}}
	route := gin.Default()
	route.POST("/build", server.BuildImage)

	payload := []byte(`{
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
    }`)

	request, err := http.NewRequest("POST", "/build", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestBuildImageShouldReturn400WhenArchitectureOfImageConfigIsMissing(t *testing.T) {
	server := &api.GinServer{ImageBuilder: &mockImageBuilder{}}
	route := gin.Default()
	route.POST("/build", server.BuildImage)

	payload := []byte(`{
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
    }`)

	request, err := http.NewRequest("POST", "/build", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
