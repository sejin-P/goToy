package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sejin-P/goToy/middlewares"
	"github.com/sejin-P/goToy/api"
)

func main() {
	r := gin.Default()

	setting := r.Group("/users")
	setting.Use(middlewares.TokenAuthMiddleware)
	{
		setting.GET("/usersetting", api.GetUserSetting)
		setting.POST("/userSetting", api.SetUserSetting)
	}

}