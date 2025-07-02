package logging

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

// InitLogger sets up zerolog to write to app.log and returns the file for closing.
func InitLogger(logPath string) (*os.File, error) {
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	zerolog.TimeFieldFormat = time.RFC3339
	Logger = zerolog.New(logFile).With().Timestamp().Logger()
	log.Logger = Logger // set global logger
	return logFile, nil
}

// ZerologMiddleware logs HTTP requests using zerolog.
func ZerologMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		event := Logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", status).
			Dur("latency", latency).
			Str("client_ip", c.ClientIP())
		if len(c.Errors) > 0 {
			event.Str("errors", c.Errors.String())
		}
		event.Msg("request completed")
	}
}
