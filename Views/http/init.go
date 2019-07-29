package http

import (
	ext "PrintHalf/Ext"
	"github.com/gin-gonic/gin"
)

var (
	route = gin.Default()
	db    = ext.GetEngine()
)

func init() {
	api := route.Group("/api")
	LoginView(api.Group("/auth"))
	//RoomView(api.Group("/room", VerifyToken))
	PictureView(api.Group("/picture", VerifyToken)) //获取图片信息
	LikeView(api.Group("/like", VerifyToken))       //点赞及查看排名
}

func GetRoute() *gin.Engine {
	return route
}
