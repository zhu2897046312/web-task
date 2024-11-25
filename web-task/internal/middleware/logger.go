package middleware

import (
    "github.com/gin-gonic/gin"
    "time"
    "log"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path

        c.Next()

        latency := time.Since(start)
        statusCode := c.Writer.Status()
        
        log.Printf("[%d] %s %s %v", statusCode, c.Request.Method, path, latency)
    }
} 