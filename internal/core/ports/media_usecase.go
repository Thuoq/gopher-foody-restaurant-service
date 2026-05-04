package ports

import (
	"context"
)

type IGetUploadURLUseCase interface {
	Execute(ctx context.Context, fileName string, contentType string) (*GeneratePresignedURLOutput, error)
}
