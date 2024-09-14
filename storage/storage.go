package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const localStorageDir = "./uploads" // Directory to store files

// Initialize the local storage directory
func InitLocalStorage() error {
	if err := os.MkdirAll(localStorageDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create local storage directory: %v", err)
	}
	return nil
}

// UploadToLocal saves the file to the local storage directory
func UploadToLocal(filename string, file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	filePath := filepath.Join(localStorageDir, filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, f)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// GenerateShareableLink returns a local file path
func GenerateShareableLink(filename string) string {
	return filepath.Join(localStorageDir, filename)
}
