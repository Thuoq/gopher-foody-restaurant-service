package dto

import (
	"time"

	"gopher-identity-service/internal/core/domain"
)

// GetProfileResponse represents the user profile data returned to the client
type GetProfileResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// MapToGetProfileResponse maps a domain.User to the GetProfileResponse DTO
func MapToGetProfileResponse(u *domain.User) GetProfileResponse {
	return GetProfileResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}
