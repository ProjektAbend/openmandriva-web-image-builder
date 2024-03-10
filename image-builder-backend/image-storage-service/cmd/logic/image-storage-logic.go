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
	destDir := "./files"

	destPath := filepath.Join(destDir, filename)
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("error: %s", err.Error())
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return fmt.Errorf("error: %s", err.Error())
	}

	return nil
}
