package websocket

import (
	config "PrintHalf/Config"
	. "PrintHalf/Models"
	"github.com/dgrijalva/jwt-go"
	"log"
)

// 校验token的函数(供socket.io使用)
func VerifyAuthToken(token string) (UserModel, string) {
	userinfo, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return config.SECRET_KEY, nil
	})
	if err != nil {
		log.Println(err.Error())
	}
	student_id := userinfo.Claims.(jwt.MapClaims)["student_id"].(string)
	user := UserModel{StudentId: student_id}
	has, err := db.Get(user)
	if err != nil {
		log.Println(err.Error())
	}
	if !has {
		return UserModel{}, "token无效"
	}
	return user, ""
}
