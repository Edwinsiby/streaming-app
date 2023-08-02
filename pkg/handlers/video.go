package handlers

import (
	"net/http"
	"stream/pkg/models"
	"stream/pkg/repository"

	"github.com/gin-gonic/gin"
)

func VideoPage(c *gin.Context) {
	c.HTML(http.StatusFound, "video.html", nil)
}

func UploadVideo(c *gin.Context) {
	var input models.Video
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.Create(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "uploaded"})
}

func StreamVideo(c *gin.Context) {
	videoFileName := c.Param("videoFileName")
	videoPath := "./video/" + videoFileName

	c.File(videoPath)
}
