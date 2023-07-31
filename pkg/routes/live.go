package routes

import (
	"stream/pkg/handlers"

	"github.com/gin-gonic/gin"
)

type LiveRouter struct{}

func NewLiveRouter() *LiveRouter {
	return &LiveRouter{}
}

func (m *LiveRouter) RegisterLiveRoutes(router *gin.RouterGroup) {

	router.GET("/", handlers.LivePage)
}
