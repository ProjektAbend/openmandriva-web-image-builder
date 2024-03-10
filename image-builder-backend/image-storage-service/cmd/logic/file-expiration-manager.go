package logic

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Metadata struct {
	TargetFile     string    `json:"filename"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
}

func CreateMetadataFile(fileName string) error {
	fileNameWithoutExtension := removeExtension(fileName)

	currentTime := time.Now()
	expirationDate := currentTime.Add(24 * time.Hour)

	metadata := Metadata{
		TargetFile:     fileName,
		ExpirationDate: expirationDate,
		CreatedAt:      currentTime,
	}

	jsonData, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return err
	}

	destPath := filepath.Join("./files", fileNameWithoutExtension+".meta.json")
	err = os.WriteFile(destPath, jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Metadata file created: %s\n", destPath)
	return nil
}

func removeExtension(filename string) string {
	lastDotIndex := strings.LastIndex(filename, ".")
	if lastDotIndex == -1 || lastDotIndex == len(filename)-1 {
		return filename
	}
	return filename[:lastDotIndex]
}
