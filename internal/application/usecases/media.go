package usecases

import (
	"context"
	"gopher-restaurant-service/internal/core/ports"
)

type mediaUseCase struct {
	storageService ports.StorageService
}

func NewMediaUseCase(storageService ports.StorageService) ports.IMediaUseCase {
	return &mediaUseCase{
		storageService: storageService,
	}
}

func (uc *mediaUseCase) GetUploadURL(ctx context.Context, fileName string, contentType string) (*ports.GeneratePresignedURLOutput, error) {
	return uc.storageService.GeneratePresignedURL(ctx, fileName, contentType)
}
