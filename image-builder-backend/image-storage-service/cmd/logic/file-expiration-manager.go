package logic

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	EXPIRATION_CHECK_PERIOD = time.Hour
	TIME_UNTIL_EXPIRATION   = 24 * time.Hour
)

type Metadata struct {
	TargetFile     string    `json:"filename"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
}

func CreateMetadataFile(fileName string) error {
	fileNameWithoutExtension := removeExtension(fileName)

	currentTime := time.Now()
	expirationDate := currentTime.Add(TIME_UNTIL_EXPIRATION)

	metadata := Metadata{
		TargetFile:     fileName,
		ExpirationDate: expirationDate,
		CreatedAt:      currentTime,
	}

	jsonData, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return err
	}

	destPath := filepath.Join(FILES_DIR, fileNameWithoutExtension+".meta.json")
	err = os.WriteFile(destPath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func removeExtension(filename string) string {
	lastDotIndex := strings.LastIndex(filename, ".")
	if lastDotIndex == -1 || lastDotIndex == len(filename)-1 {
		return filename
	}
	return filename[:lastDotIndex]
}

func DeleteExpiredFilesAndMetaData() {
	for {
		log.Print("Checking for expired files...")
		currentTime := time.Now()

		files, err := os.ReadDir(FILES_DIR)
		if err != nil {
			log.Printf("Error reading directory: %s", err)
			continue
		}

		for _, file := range files {
			if !strings.HasSuffix(file.Name(), ".meta.json") {
				continue
			}

			metaFilePath := filepath.Join(FILES_DIR, file.Name())
			metaFileContent, err := os.ReadFile(metaFilePath)
			if err != nil {
				log.Printf("Error reading metadata file: %s", err)
				continue
			}

			var metadata Metadata
			if err := json.Unmarshal(metaFileContent, &metadata); err != nil {
				log.Printf("Error parsing metadata JSON: %s", err)
				continue
			}

			if currentTime.After(metadata.ExpirationDate) {
				filePath := filepath.Join(FILES_DIR, metadata.TargetFile)
				if err := os.Remove(filePath); err != nil {
					log.Printf("Error deleting file: %s", err)
					continue
				}
				log.Printf("Deleted file: %s", filePath)

				if err := os.Remove(metaFilePath); err != nil {
					log.Printf("Error deleting metadata file %s: %s", metaFilePath, err)
					continue
				}
				log.Printf("Deleted metadata file: %s\n", metaFilePath)
			}
		}
		time.Sleep(EXPIRATION_CHECK_PERIOD)
	}
}
