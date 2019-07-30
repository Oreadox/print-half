package websocket

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
)

func Connect(s socketio.Conn) error {
	log.Println("on connection")
	s.Emit("connect", jsonify{
		"message": "连接建立成功",
	})
	fmt.Printf("%+v", s)
	return nil
}

func DisConnect(s socketio.Conn, msg string) {
	fmt.Println("closed", msg)
	j := 0
	for _, val := range matching {
		if val != s {
			matching[j] = val
			j++
		}
	}
	matching = matching[:j]
}
