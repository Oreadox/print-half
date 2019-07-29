package http

import (
	. "PrintHalf/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginView(g *gin.RouterGroup) {
	g.POST("main", func(c *gin.Context) {
		if result, status, err := MainLogin(c); err != nil {
			c.AbortWithStatusJSON(status, result)
		} else {
			c.JSON(status, result)
		}
	})
}

func MainLogin(c *gin.Context) (*map[string]interface{}, int, error) {
	var userinfo LoginModel
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusBadRequest, err
	}
	user := UserModel{
		StudentId: userinfo.StudentId,
		Name:      userinfo.Name,
	}
	has, err := db.Get(&user)
	if err != nil {
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusBadRequest, err
	}
	var token string
	if !has {
		_, err := db.Insert(user)
		if err != nil {
			return &map[string]interface{}{
				"message": err.Error(),
				"status":  0,
			}, http.StatusBadRequest, err
		}
	}
	token = userinfo.GenerateToken()
	return &map[string]interface{}{
		"message": "登录成功！",
		"status":  1,
		"data": map[string]interface{}{
			"token": token,
		},
	}, http.StatusOK, nil
}
