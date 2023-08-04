package music

import (
	"github.com/gin-gonic/gin"
)

type MusicRouter struct{}

func NewMusicRouter() *MusicRouter {
	return &MusicRouter{}
}

func (m *MusicRouter) RegisterMusicRoutes(router *gin.RouterGroup) {
	router.GET("/", MusicPage)
	router.POST("/upload/image", UploadMusic)
	router.GET("/:musicFileName", StreamMusic)
}
