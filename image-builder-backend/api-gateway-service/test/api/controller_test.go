package api

import (
	"bytes"
	"encoding/json"
	"github.com/api-gateway-service/cmd/api"
	"github.com/gin-gonic/gin"
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

func TestBuildImageShouldAlsoReturnValidImageInfoWhenReturning201(t *testing.T) {
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

	var responseImageInfo api.ImageInfo
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseImageInfo); err != nil {
		t.Errorf("Failed to parse response JSON: %v", err)
		return
	}

	expectedImageInfo2 := `{
		"availableUntil": null,
		"imageId":        "a1b2c3",
		"isAvailable":    null
	}`

	actualImageInfo := recorder.Body.String()

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
	server := &api.GinServer{ImageBuilder: &mockImageBuilder{}}
	route := gin.Default()
	route.POST("/build", server.BuildImage)

	payload := []byte(`{
        "architecture": "aarch64-uefi"
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
