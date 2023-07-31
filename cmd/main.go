package main

import (
	"stream/pkg/handlers"
	"stream/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := NewRouter()
	router.Run(":8080")
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", handlers.IndexPage)

	musicGroup := router.Group("/music")
	musicRouter := routes.NewMusicRouter()
	musicRouter.RegisterMusicRoutes(musicGroup)

	videoGroup := router.Group("/video")
	videoRouter := routes.NewVideoRouter()
	videoRouter.RegisterVideoRoutes(videoGroup)

	return router
}
