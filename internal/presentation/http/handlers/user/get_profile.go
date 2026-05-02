package user

import (
	responseDto "gopher-identity-service/internal/presentation/http/handlers/user/dto/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"gopher-identity-service/internal/core/ports"
	"gopher-identity-service/pkg/response"
)

type GetProfileHandler struct {
	ssoUseCase ports.SSOUseCase
	logger     *zap.Logger
}

func NewGetProfileHandler(ssoUseCase ports.SSOUseCase, logger *zap.Logger) *GetProfileHandler {
	return &GetProfileHandler{
		ssoUseCase: ssoUseCase,
		logger:     logger,
	}
}

func (h *GetProfileHandler) Handle(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.logger.Error("Invalid user ID format", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.ssoUseCase.GetUserProfile(c.Request.Context(), userID)
	if err != nil {
		h.logger.Error("Failed to get user profile", zap.Error(err), zap.Int64("user_id", userID))
		response.Error(c, http.StatusInternalServerError, "internal server error or user not found")
		return
	}

	res := responseDto.MapToGetProfileResponse(user)
	response.Success(c, http.StatusOK, res)
}
