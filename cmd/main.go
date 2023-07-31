package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", indexPage)

	router.GET("/music", musicPage)
	router.POST("/upload/image", uploadMusic)
	router.GET("/music/:musicFileName", streamMusic)

	router.GET("/video", videoPage)
	router.POST("/upload/video", uploadVideo)
	router.GET("/video/:videoFileName", streamVideo)

	router.Run(":8080")
}

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func musicPage(c *gin.Context) {
	c.HTML(http.StatusFound, "music.html", nil)
}

func uploadMusic(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "uploaded"})
}

func streamMusic(c *gin.Context) {
	musicFileName := c.Param("musicFileName")
	musicPath := "./music/" + musicFileName
	fmt.Println(musicPath)
	c.File(musicPath)
}

func videoPage(c *gin.Context) {
	c.HTML(http.StatusFound, "video.html", nil)
}

func uploadVideo(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "uploaded"})
}

func streamVideo(c *gin.Context) {
	videoFileName := c.Param("videoFileName")
	videoPath := "./video/" + videoFileName

	c.File(videoPath)
}
