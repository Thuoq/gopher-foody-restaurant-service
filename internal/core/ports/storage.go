package ports

import (
	"context"
)

type GeneratePresignedURLOutput struct {
	UploadURL string `json:"upload_url"`
	FinalURL  string `json:"final_url"`
}

type StorageService interface {
	GeneratePresignedURL(ctx context.Context, fileName string, contentType string) (*GeneratePresignedURLOutput, error)
}
