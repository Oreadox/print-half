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
	RoomView(api.Group("/room", VerifyToken))
}

func GetRoute() *gin.Engine {
	return route
}
