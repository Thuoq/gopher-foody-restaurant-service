package usecases

import (
	"context"

	"gopher-identity-service/internal/core/domain"
	"gopher-identity-service/internal/core/ports"
)

type ssoUseCase struct {
	userRepo ports.UserRepository
}

func NewSSOUseCase(userRepo ports.UserRepository) ports.SSOUseCase {
	return &ssoUseCase{
		userRepo: userRepo,
	}
}

func (uc *ssoUseCase) GetUserProfile(ctx context.Context, userID int64) (*domain.User, error) {
	// Add business logic here if needed (e.g. check if user is active, mask data, etc)
	return uc.userRepo.GetByID(ctx, userID)
}
