package user

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	getProfileHandler *GetProfileHandler
}

func NewRouter(getProfileHandler *GetProfileHandler) *Router {
	return &Router{
		getProfileHandler: getProfileHandler,
	}
}

func (r *Router) Register(api *gin.RouterGroup) {
	userGroup := api.Group("/sso")
	{
		userGroup.GET("/profile/:id", r.getProfileHandler.Handle)
	}
}
