package video

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func VideoPage(c *gin.Context) {
	c.HTML(http.StatusFound, "video.html", nil)
}

func UploadVideo(c *gin.Context) {
	var input Video
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := Create(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "uploaded"})
}

func StreamVideo(c *gin.Context) {
	sectionID := c.Param("sectionID")

	sectionIDInt, err := strconv.Atoi(sectionID)
	if err != nil {
		fmt.Println("error", err)
	}

	videoFileName, err := FindBySectionID(sectionIDInt)
	if err != nil {
		fmt.Println("Video not found", err)
	}
	videoPath := "./static/video/" + videoFileName

	c.File(videoPath)
}
