package video

import (
	"github.com/gin-gonic/gin"
)

type VideoRouter struct{}

func NewVideoRouter() *VideoRouter {
	return &VideoRouter{}
}

func (m *VideoRouter) RegisterVideoRoutes(router *gin.RouterGroup) {

	router.GET("/", VideoPage)
	router.POST("/upload/video", UploadVideo)
	router.GET("/:sectionID", StreamVideo)
	router.GET("/hls", HlsVideoPage)
	router.GET("/hls/:sectionID", StreamVideoHls)
}
