package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("stream.html")
	})

	r.GET("/stream", streamHandler)
	r.GET("/download", downloadHandler)

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

func downloadHandler(c *gin.Context) {
	videoID := c.Query("id")
	if videoID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing video ID or URL"})
		return
	}

	videoID = extractVideoID(videoID)
	if videoID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid video ID or URL"})
		return
	}

	cmd := exec.Command("yt-dlp", "-f", "best", "-o", fmt.Sprintf("%s.mp4", videoID), "https://www.youtube.com/watch?v="+videoID)

	err := cmd.Start()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error starting command: %v", err)})
		return
	}

	defer func() {
		if cmd.ProcessState == nil || !cmd.ProcessState.Exited() {
			cmd.Process.Kill()
		}
	}()
	cmd.Wait()
	time.Sleep(time.Second)
}

func streamHandler(c *gin.Context) {
	videoID := c.Query("id")
	if videoID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing video ID or URL"})
		return
	}

	videoID = extractVideoID(videoID)
	if videoID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid video ID or URL"})
		return
	}

	cmd := exec.Command("yt-dlp", "-f", "best", "-o", "-", "https://www.youtube.com/watch?v="+videoID)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error creating stdout pipe: %v", err)})
		return
	}

	err = cmd.Start()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error starting command: %v", err)})
		return
	}

	defer func() {
		if cmd.ProcessState == nil || !cmd.ProcessState.Exited() {
			cmd.Process.Kill()
		}
	}()

	c.Header("Content-Type", "video/mp4")
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s.mp4", videoID))

	// Stream the video to the client
	_, err = io.Copy(c.Writer, stdout)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error streaming video: %v", err)})
		return
	}

	// Wait for the command to finish (not strictly necessary, but good practice)
	cmd.Wait()

	// Sleep for a short while to allow the client to finish reading the response
	time.Sleep(time.Second)
}

func extractVideoID(idOrURL string) string {
	// First, try to match the input against the YouTube video ID pattern
	match := regexp.MustCompile(`^[a-zA-Z0-9_-]{11,64}$`).FindString(idOrURL)
	if match != "" {
		return match
	}

	// If it doesn't match, try to parse the input as a YouTube video URL
	if !strings.HasPrefix(idOrURL, "http") {
		idOrURL = "http://" + idOrURL
	}
	u, err := url.Parse(idOrURL)
	if err == nil && u.Host == "www.youtube.com" || u.Host == "youtube.com" && u.Path == "/watch" {
		q := u.Query()
		return q.Get("v")
	}

	return ""
}
