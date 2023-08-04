package live

import (
	"github.com/gin-gonic/gin"
)

func StreamCamRoutes(router *gin.RouterGroup, handler *StreamCamHandler) {
	router.GET("/", handler.LivePage)
	router.GET("/cam", handler.LiveStreamCam)
}
