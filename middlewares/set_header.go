package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
)

func SetHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
	c.Header("Last-modified", time.Now().String())
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "-1")
	c.Next()
}
