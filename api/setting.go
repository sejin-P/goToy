package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sejin-P/goToy/database"
	"log"
	"net/http"
)

type UserInput struct {
	Email string
}

func GetUserSetting(ctx *gin.Context) {
	var input UserInput

	err := ctx.Bind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := User{}
	user.Email = input.Email
	db, err := database.ConnectToDB()
	defer db.Close()
	if err != nil {
		log.Println(err)
		return
	}
	_ = db.QueryRow("SELECT user_no, user_name, user_email FROM user_info WHERE user_id = $1 AND user_pw = $2 AND is_enabled = 1",
		).Scan(&user.UserNo, &user.Name, &user.Email)

}
