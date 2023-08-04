package main

import (
	"log"
	"stream/pkg/live"
	"stream/pkg/music"
	"stream/pkg/video"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	router := NewRouter()
	router.Run(":8080")
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", music.IndexPage)

	musicGroup := router.Group("/music")
	musicRouter := music.NewMusicRouter()
	musicRouter.RegisterMusicRoutes(musicGroup)

	videoGroup := router.Group("/video")
	videoRouter := video.NewVideoRouter()
	videoRouter.RegisterVideoRoutes(videoGroup)

	err := InitializeStreamServer(router)
	if err != nil {
		log.Fatalln(err)
	}
	return router
}

func InitializeStreamServer(router *gin.Engine) error {

	StreamCamConn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		return err
	}
	StreamCamHandler, err := initializeAuthenticationHandler(StreamCamConn)
	if err != nil {
		return err
	}
	live.StreamCamRoutes(router.Group("/live"), StreamCamHandler)
	return nil
}

func initializeAuthenticationHandler(cc *grpc.ClientConn) (*live.StreamCamHandler, error) {
	return live.NewAuthenticationHandler(cc), nil
}
