package routes

import (
	"stream/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func StreamCamRoutes(router *gin.RouterGroup, handler *handlers.StreamCamHandler) {
	router.GET("/", handler.LivePage)
	router.GET("/cam", handler.LiveStreamCam)
}
