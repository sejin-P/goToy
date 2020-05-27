package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sejin-P/goToy/api"
	"net/http"
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
			"error":  "token is None",
		})
		return
	}
	claims := &api.Claims{}

	_, err = jwt.ParseWithClaims(tkn, claims, func(token *jwt.Token) (interface{}, error) {
		return api.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "token is expired"})
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "error": "auth failed"})
		return
	} else {
		c.Next()
	}
}
