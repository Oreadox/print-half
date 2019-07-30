package websocket

import (
	. "PrintHalf/Models"
	"github.com/googollee/go-socket.io"
)

var matching []socketio.Conn

func Join(s socketio.Conn, json map[string]interface{}) {
	token := json["token"].(string)
	user, err := VerifyAuthToken(token)
	if err != "" {
		s.Emit("join", jsonify{
			"message": err,
		})
	} else {
		s.SetContext(user.Id)
		s.Emit("join", jsonify{
			"message": "进入房间成功",
		})
		if len(matching) != 1 {
			match(s, matching[0])
		} else {
			matching = append(matching, s)
		}
	}
}

func Exit(s socketio.Conn) {
	j := 0
	for _, val := range matching {
		if val != s {
			matching[j] = val
			j++
		}
	}
	matching = matching[:j]
	s.Emit("exit", jsonify{
		"message": "退出房间成功",
	})
}

// 匹配
func match(s1, s2 socketio.Conn) {
	if s1 == s2 {
		return
	} else {
		picture := PictureModel{
			UserId1: s1.Context().(int),
			UserId2: s2.Context().(int),
		}
		db.Insert(picture)
		matching = matching[1:]
		s1.Emit("match", jsonify{
			"message":         "匹配成功",
			"another_user_id": s2.Context().(int),
			// 其他再加
		})
		s2.Emit("match", jsonify{
			"message":         "匹配成功",
			"another_user_id": s1.Context().(int),
			// 其他再加
		})
	}
}
