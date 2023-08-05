package video

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func VideoPage(c *gin.Context) {
	c.HTML(http.StatusFound, "video.html", nil)
}

func HlsVideoPage(c *gin.Context) {
	c.HTML(http.StatusFound, "hls.html", nil)
}

func UploadVideo(c *gin.Context) {
	var input Video
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	videoFileName, err := FindBySectionID(input.Section)
	if err != nil {
		fmt.Println("Video not found", err)
	}
	if videoFileName != "" {
		if err := Update(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
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

func StreamVideoHls(c *gin.Context) {
	// buffer := make([]byte, 10)
	sectionID := c.Param("sectionID")

	sectionIDInt, err := strconv.Atoi(sectionID)
	if err != nil {
		fmt.Println("error", err)
		c.Status(http.StatusBadRequest)
		return
	}

	videoFileName, err := FindBySectionID(sectionIDInt)
	if err != nil {
		fmt.Println("Video not found", err)
		c.Status(http.StatusNotFound)
		return
	}
	videoPath := "./static/video/" + videoFileName

	hlsOutputPath := "./static/hls_output"

	err = videoToChunks(videoPath, hlsOutputPath, "output", 10)
	if err != nil {
		fmt.Println("Error generating HLS files:", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	hlsManifestPath := hlsOutputPath + "/output.m3u8"

	playlist, err := os.Open(hlsManifestPath)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		fmt.Println(err)
		return
	}
	defer playlist.Close()

	c.Header("Content-Type", "application/vnd.apple.mpegurl")
	c.Header("Content-Disposition", "inline")

	playlist.Seek(0, 0)

	_, err = io.Copy(c.Writer, playlist)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		fmt.Println(err)
	}
}

func videoToChunks(inputVideoPath, outputDir, outputFilename string, chunkDuration int) error {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return err
		}
	}

	segmentPattern := filepath.Join(outputDir, outputFilename+"_%%03d.ts")
	cmd := exec.Command(
		"ffmpeg",
		"-i", inputVideoPath,
		"-c:v", "copy",
		"-c:a", "libfdk_aac",
		"-hls_time", strconv.Itoa(chunkDuration),
		"-hls_list_size", "0",
		"-hls_segment_filename", segmentPattern,
		"-strftime", "1",
		filepath.Join(outputDir, outputFilename+".m3u8"),
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return err
	}

	return nil
}
