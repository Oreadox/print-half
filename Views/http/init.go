package http

import (
	ext "PrintHalf/Ext"
	"PrintHalf/Utils"
	"github.com/gin-gonic/gin"
)

var (
	route = gin.Default()
	db    = ext.GetEngine()
)

func init() {
	api := route.Group("/api")
	LoginView(api.Group("/auth"))
	RoomView(api.Group("/room", utils.VerifyToken))
	PictureView(api.Group("/picture"))              //获取图片信息
	LikeView(api.Group("/like", utils.VerifyToken)) //点赞及查看排名
}

func GetRoute() *gin.Engine {
	return route
}
