package websocket

import (
	. "PrintHalf/Models"
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
)

func Connect(s socketio.Conn) error {
	log.Println("on connection")
	onlineNum := StatusModel{Desc: "在线人数"}
	_, err := db.Get(&onlineNum)
	onlineNum.Value += 1
	_, err = db.Id(onlineNum.Id).Cols("value").Update(&onlineNum)
	if err != nil {
		s.Close()
		return err
	}
	s.Emit("connect", jsonify{
		"message": "连接建立成功",
	})
	fmt.Printf("%+v", s)
	return nil
}

func DisConnect(s socketio.Conn, msg string) {
	fmt.Println("closed", msg)
	onlineNum := StatusModel{Desc: "在线人数"}
	_, err := db.Get(&onlineNum)
	defer func(err error) {
		if err != nil {
			log.Println(err.Error())
		}
	}(err)
	onlineNum.Value -= 1
	_, err = db.Id(onlineNum.Id).Cols("value").Update(&onlineNum)
}
