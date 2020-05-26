package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sejin-P/goToy/middlewares"
)

func main() {
	r := gin.Default()

	setting := r.Group("/users")
	setting.Use(middlewares.TokenAuthMiddleware)

}