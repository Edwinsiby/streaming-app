package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LivePage(c *gin.Context) {
	c.HTML(http.StatusFound, "live.html", nil)
}

func LiveStreamCam(c *gin.Context) {

}
