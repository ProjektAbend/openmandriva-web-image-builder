package logic

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type ImageStorageLogic struct{}

func (c *ImageStorageLogic) StoreImage(file multipart.File, filename string) error {
	destPath := filepath.Join("./files", filename)
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err.Error())
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return fmt.Errorf("error copying contents to file: %s", err.Error())
	}

	return nil
}

func (c *ImageStorageLogic) GetIsoFile(fileName string) (string, error) {
	return filepath.Join("./files", fileName), nil
}

func (c *ImageStorageLogic) DoesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
