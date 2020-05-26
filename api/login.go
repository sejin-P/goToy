package api

import (
	"crypto/sha512"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	db2 "github.com/sejin-P/goToy/db"
	"log"
	"time"
)

type User struct {
	Email string
	Pw string
	UserNo int
	Name string
}

type Claims struct {
	UserNo int
	jwt.StandardClaims
}

var expirationTime = 5*time.Minute

var JwtKey = []byte("signed")

func (user User) LoginUser() bool {
	db, err := db2.ConnectToDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	_ = db.QueryRow("SELECT user_no, user_name FROM user_info WHERE user_id = $1 AND user_pw = $2 AND is_enabled = 1",
		user.Email, sha512.Sum512([]byte(user.Pw))).Scan(&user.UserNo, &user.Name)

	if user.UserNo != 0 {
		return true
	} else {
		return false
	}
}

func (user *User) GetJwtToken() (string, error) {
	expirationTime := time.Now().Add(expirationTime)
	claims := &Claims{
		UserNo: user.UserNo,
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
