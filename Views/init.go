package views

import (
	"PrintHalf/Views/http"
	"PrintHalf/Views/websocket"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
)

var (
	route  = http.GetRoute()
	socket = websocket.GetSocket()
)

func GetRoute() *gin.Engine {
	return route
}

func GetSocket() *socketio.Server {
	return socket
}
