package admin

import (
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	mediaUseCase ports.IMediaUseCase
}

func NewMediaHandler(mediaUseCase ports.IMediaUseCase) *MediaHandler {
	return &MediaHandler{
		mediaUseCase: mediaUseCase,
	}
}

type uploadURLRequest struct {
	FileName    string `json:"file_name" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
}

func (h *MediaHandler) GetUploadURL(c *gin.Context) {
	var req uploadURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.mediaUseCase.GetUploadURL(c.Request.Context(), req.FileName, req.ContentType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to generate upload url")
		return
	}

	response.Success(c, http.StatusOK, result)
}
