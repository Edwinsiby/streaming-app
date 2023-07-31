package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideoPage(c *gin.Context) {
	c.HTML(http.StatusFound, "video.html", nil)
}

func UploadVideo(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "uploaded"})
}

func StreamVideo(c *gin.Context) {
	videoFileName := c.Param("videoFileName")
	videoPath := "./video/" + videoFileName

	c.File(videoPath)
}
