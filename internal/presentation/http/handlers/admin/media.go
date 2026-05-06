package admin

import (
	"gopher-restaurant-service/internal/core/ports"
	"gopher-restaurant-service/pkg/app_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	getUploadURLUC ports.IGetUploadURLUseCase
}

func NewMediaHandler(getUploadURLUC ports.IGetUploadURLUseCase) *MediaHandler {
	return &MediaHandler{
		getUploadURLUC: getUploadURLUC,
	}
}

type uploadURLRequest struct {
	FileName    string `json:"file_name" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
}

func (h *MediaHandler) GetUploadURL(c *gin.Context) {
	var req uploadURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fieldErrors := app_response.ParseValidationErrors(err)
		app_response.ValidationError(c, http.StatusBadRequest, "invalid request body", fieldErrors)
		return
	}

	result, err := h.getUploadURLUC.Execute(c.Request.Context(), req.FileName, req.ContentType)
	if err != nil {
		app_response.Error(c, http.StatusInternalServerError, "failed to generate upload url")
		return
	}

	app_response.Success(c, http.StatusOK, result)
}
