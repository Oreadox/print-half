package websocket

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
)

var userBroadcast = socketio.NewBroadcast()

func Connect(s socketio.Conn) error {
	playerBroadcast.Join("allPlayer", s)
	log.Println("on connection")
	s.Emit("connect", jsonify{
		"message": "连接建立成功",
	})
	fmt.Printf("%+v", s)
	//s.c
	return nil
}

func DisConnect(s socketio.Conn, msg string) {
	fmt.Println("closed", msg)
	j := 0
	for _, val := range matching {
		if val.Sid != s.ID() {
			matching[j] = val
			j++
		}
	}
	matching = matching[:j]
}
