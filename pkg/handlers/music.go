package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func MusicPage(c *gin.Context) {
	c.HTML(http.StatusFound, "music.html", nil)
}

func UploadMusic(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "uploaded"})
}

func StreamMusic(c *gin.Context) {
	musicFileName := c.Param("musicFileName")
	musicPath := "./music/" + musicFileName
	fmt.Println(musicPath)
	c.File(musicPath)
}
