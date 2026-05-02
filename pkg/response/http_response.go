package response

import "github.com/gin-gonic/gin"

// Response represents a standard API response body
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success writes a standard success response
func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: "success",
		Data:    data,
	})
}

// Error writes a standard error response
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}
