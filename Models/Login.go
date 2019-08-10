package models

import (
	"PrintHalf/Config"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type LoginModel struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	jwt.StandardClaims
}

func (user *LoginModel) GenerateToken() string {
	user.ExpiresAt = time.Now().Add(time.Hour * config.ExpiresTime).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	token, err := t.SignedString([]byte(config.SecretKey))
	if err != nil {
		log.Println(err.Error())
	}
	return token
}

//
//func (user *LoginModel)VerifyToken(token string) *UserModel {
//
//}
