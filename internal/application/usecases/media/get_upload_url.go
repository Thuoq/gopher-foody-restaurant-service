package media

import (
	"context"
	"gopher-restaurant-service/internal/core/ports"
)

type getUploadURLUseCase struct {
	storageService ports.StorageService
}

func NewGetUploadURLUseCase(storageService ports.StorageService) ports.IGetUploadURLUseCase {
	return &getUploadURLUseCase{
		storageService: storageService,
	}
}

func (uc *getUploadURLUseCase) Execute(ctx context.Context, fileName string, contentType string) (*ports.GeneratePresignedURLOutput, error) {
	return uc.storageService.GeneratePresignedURL(ctx, fileName, contentType)
}
