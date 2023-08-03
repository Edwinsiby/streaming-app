package routes

import (
	"stream/pkg/handlers"

	"github.com/gin-gonic/gin"
)

type VideoRouter struct{}

func NewVideoRouter() *VideoRouter {
	return &VideoRouter{}
}

func (m *VideoRouter) RegisterVideoRoutes(router *gin.RouterGroup) {

	router.GET("/", handlers.VideoPage)
	router.POST("/upload/video", handlers.UploadVideo)
	router.GET("/:sectionID", handlers.StreamVideo)
}
