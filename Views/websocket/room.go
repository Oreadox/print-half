package websocket

import (
	. "PrintHalf/Models"
	utils "PrintHalf/Utils"
	"github.com/googollee/go-socket.io"
	"log"
	"math/rand"
	"time"
	"unsafe"
)

var matching []socketio.Conn

func Join(s socketio.Conn, json map[string]interface{}) {
	token := json["token"].(string)
	user, err := utils.VerifyAuthToken(token)
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
		count, err := db.Count(&QuestionModel{})
		var question *QuestionModel
		var has bool
		for !has {
			rand.Seed(time.Now().UnixNano())
			num := rand.Intn(*(*int)(unsafe.Pointer(&count))) + 1 //[1,count]
			question.Id = num
			has, err = db.Get(num)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
		picture := PictureModel{
			UserId1:  s1.Context().(int),
			UserId2:  s2.Context().(int),
			Question: question.Name,
		}
		user1 := UserModel{Id: picture.UserId1}
		user2 := UserModel{Id: picture.UserId2}
		_, err = db.Get(user1)
		_, err = db.Get(user2)
		if err != nil {
			log.Println(err.Error())
			return
		}
		db.Insert(picture)
		matching = matching[1:]
		s1.Emit("match", jsonify{
			"message": "匹配成功",
			"data": jsonify{
				"another_user_name": user2.Name,
				"question":          question.Name,
			},
			// 其他再加
		})
		s2.Emit("match", jsonify{
			"message": "匹配成功",
			"data": jsonify{
				"another_user_name": user1.Name,
				"question":          question.Name,
			},
			// 其他再加
		})
	}
}
