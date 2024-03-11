package mocks

import (
	"fmt"
	"mime/multipart"
)

type MockImageStorageLogic struct{}

func (_ *MockImageStorageLogic) StoreImage(_ multipart.File, _ string) error {
	return nil
}

func (_ *MockImageStorageLogic) GetIsoFile(_ string) (string, error) {
	return "", nil
}

func (_ *MockImageStorageLogic) DoesFileExist(_ string) bool {
	return true
}

type MockImageStorageLogicReturnsError struct{}

func (_ *MockImageStorageLogicReturnsError) StoreImage(_ multipart.File, _ string) error {
	return fmt.Errorf("error occurred")
}

func (_ *MockImageStorageLogicReturnsError) GetIsoFile(_ string) (string, error) {
	return "", fmt.Errorf("error occurred")
}

func (_ *MockImageStorageLogicReturnsError) DoesFileExist(_ string) bool {
	return true
}
