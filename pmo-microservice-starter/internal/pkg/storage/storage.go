package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type StorageService struct {
	basePath string
	baseURL  string
}

func NewStorageService(basePath, baseURL string) *StorageService {

	// Create directory if not exists
	os.MkdirAll(basePath, os.ModePerm)
	return &StorageService{
		basePath: basePath,
		baseURL:  baseURL,
	}
}

func (s *StorageService) SaveFile(file *multipart.FileHeader, subDir string) (string, int64, error) {

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), strings.TrimSuffix(file.Filename, ext), ext)

	// Create subdirectory
	fullPath := filepath.Join(s.basePath, subDir)
	os.MkdirAll(fullPath, os.ModePerm)

	// Full file path
	filePath := filepath.Join(fullPath, filename)

	// Save file
	src, err := file.Open()
	if err != nil {
		return "", 0, err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return "", 0, err
	}
	defer dst.Close()

	// Copy file
	bytes, err := io.Copy(dst, src)
	if err != nil {
		return "", 0, err
	}
	// Generate URL
	fileURL := fmt.Sprintf("%s/%s/%s", s.baseURL, subDir, filename)
	return fileURL, bytes, nil
}
func (s *StorageService) DeleteFile(fileURL string) error {
	// Extract path from URL
	if strings.HasPrefix(fileURL, s.baseURL) {
		relativePath := strings.TrimPrefix(fileURL, s.baseURL)
		filePath := filepath.Join(s.basePath, relativePath)
		return os.Remove(filePath)
	}
	return nil
}
