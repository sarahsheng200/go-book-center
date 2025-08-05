package middleware

import (
	"github.com/gin-gonic/gin"
	"go-book-center/app/common"
	"log"
	"time"
)

var logger = common.Logger

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		log.Printf("LoggerToFile:  %v |", c)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latency, clientIP, reqMethod, reqUri)
	}
}
