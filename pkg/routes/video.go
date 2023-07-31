package routes

import (
	"stream/pkg/handlers"

	"github.com/gin-gonic/gin"
)

type VideoRouter struct{}

func NewVideoRouter() *MusicRouter {
	return &MusicRouter{}
}

func (m *MusicRouter) RegisterVideoRoutes(router *gin.RouterGroup) {

	router.GET("/", handlers.VideoPage)
	router.POST("/upload/video", handlers.UploadVideo)
	router.GET("/:videoFileName", handlers.StreamVideo)
}
