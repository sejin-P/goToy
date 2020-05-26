package middlewares

import (
	"github.com/sejin-P/goToy/api"
	"log"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware(c *gin.Context) {
	token, err := c.Request.Cookie("access-token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Authentication failed"})
		return
	}

	tkn := token.Value
	if tkn == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error": "token is None",
		})
		return
	}
	claims := &api.Claims{}

	_, err = jwt

}