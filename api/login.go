package api

import (
	"crypto/sha512"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sejin-P/goToy/database"
	"github.com/pkg/errors"
	"log"
	"time"
)

type User struct {
	Email  string
	Pw     string
	UserNo int
	Name   string
}

type Claims struct {
	UserNo int
	jwt.StandardClaims
}

var expirationTime = 5 * time.Minute

var JwtKey = []byte("signed")

func (user User) LoginUser() (bool, error) {
	db, err := database.ConnectToDB()
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false, errors.Wrapf(err, "db cannot connect")
	}

	_ = db.QueryRow("SELECT user_no, user_name FROM user_info WHERE user_id = $1 AND user_pw = $2 AND is_enabled = 1",
		user.Email, sha512.Sum512([]byte(user.Pw))).Scan(&user.UserNo, &user.Name)

	if user.UserNo != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (user *User) GetJwtToken() (string, error) {
	expirationTime := time.Now().Add(expirationTime)
	claims := &Claims{
		UserNo:         user.UserNo,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", fmt.Errorf("token signed error")
	} else {
		return tokenString, nil
	}
}
