package http

import (
	"github.com/gin-gonic/gin"
)

func RoomView(g *gin.RouterGroup) {
	g.POST("", func(c *gin.Context) {
		if result, status, err := Login(c); err != nil {
			c.AbortWithStatusJSON(status, result)
		} else {
			c.JSON(status, result)
		}
	})
}

//
//func JoinRoom(c gin.Context) (*map[string]interface{}, int, error) {
//	return nil,0, nil
//}
