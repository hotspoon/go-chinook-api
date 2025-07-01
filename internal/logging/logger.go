package logging

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"io"

	"github.com/gin-gonic/gin"
)

// InitLogger sets up logging to both stdout and a file.
func InitLogger(logPath string) (*os.File, error) {
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	return logFile, nil
}

// JSONLogger returns a Gin middleware that logs requests in JSON format.
func JSONLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		entry := map[string]interface{}{
			"time":       time.Now().Format(time.RFC3339),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     status,
			"latency_ms": latency.Milliseconds(),
			"client_ip":  c.ClientIP(),
		}
		if len(c.Errors) > 0 {
			entry["errors"] = c.Errors.String()
		}
		logLine, _ := json.Marshal(entry)
		log.Println(string(logLine))
	}
}
