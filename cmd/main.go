package main

import (
	"log"
	"stream/pkg/handlers"
	"stream/pkg/routes"

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

	router.GET("/", handlers.IndexPage)

	musicGroup := router.Group("/music")
	musicRouter := routes.NewMusicRouter()
	musicRouter.RegisterMusicRoutes(musicGroup)

	videoGroup := router.Group("/video")
	videoRouter := routes.NewVideoRouter()
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
	routes.StreamCamRoutes(router.Group("/live"), StreamCamHandler)
	return nil
}

func initializeAuthenticationHandler(cc *grpc.ClientConn) (*handlers.StreamCamHandler, error) {
	return handlers.NewAuthenticationHandler(cc), nil
}
