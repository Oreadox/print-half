package utils

import (
	config "PrintHalf/Config"
	ext "PrintHalf/Ext"
	. "PrintHalf/Models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var db = ext.GetEngine()

// 校验token的中间件
func VerifyToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	userinfo, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
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
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "token无效",
		})
	}
	c.Set("user", user)
	c.Next()
}

// 校验token的函数(供socket.io使用)
func VerifyAuthToken(token string) (UserModel, string) {
	userinfo, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
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
