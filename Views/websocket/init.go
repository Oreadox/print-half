package websocket

import (
	ext "PrintHalf/Ext"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

type jsonify map[string]interface{}

var (
	db     = ext.GetEngine()
	socket *socketio.Server
)

func init() {
	var err error
	socket, err = socketio.NewServer(nil)
	if err != nil {
		log.Println(err.Error())
	}
	socketInit()
}

func socketInit() {
	socket.OnConnect("/", Connect)
	socket.OnDisconnect("/", DisConnect)
	socket.OnError("/", func(e error) {
		log.Println("meet error:", e)
	})
	socket.OnEvent("/", "join", Join)
}

func GetSocket() *socketio.Server {
	return socket
}
