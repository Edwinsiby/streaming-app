package routes

import (
	"stream/pkg/handlers"

	"github.com/gin-gonic/gin"
)

type MusicRouter struct{}

func NewMusicRouter() *MusicRouter {
	return &MusicRouter{}
}

func (m *MusicRouter) RegisterMusicRoutes(router *gin.RouterGroup) {
	router.GET("/", handlers.MusicPage)
	router.POST("/upload/image", handlers.UploadMusic)
	router.GET("/:musicFileName", handlers.StreamMusic)
}
