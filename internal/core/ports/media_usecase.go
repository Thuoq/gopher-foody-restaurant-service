package ports

import (
	"context"
)

type IMediaUseCase interface {
	GetUploadURL(ctx context.Context, fileName string, contentType string) (*GeneratePresignedURLOutput, error)
}
