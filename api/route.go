package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInput struct {
	Email string
	Pw    string
}

func GetLoginUser(ctx *gin.Context) {
	var input LoginInput
	err := ctx.Bind(&input)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, "cannot get user info")
		return
	}



}
