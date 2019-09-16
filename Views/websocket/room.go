package websocket

import (
	. "PrintHalf/Models"
	utils "PrintHalf/Utils"
	"fmt"
	"github.com/googollee/go-socket.io"
	jsoniter "github.com/json-iterator/go"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type MatchingModel struct {
	Sid    string
	UserId int
}

var matching []MatchingModel
var broadcast = socketio.NewBroadcast()

func Join(s socketio.Conn, data string) {

	var tokenData struct {
		Token string
	}
	broadcast.Join(s.ID(), s)
	jsoniter.Unmarshal([]byte(data), &tokenData)
	token := tokenData.Token
	fmt.Println("token:" + token)
	//fmt.Println(token)
	user, err := utils.VerifyAuthToken(token)

	if err != "" {
		s.Emit("join", jsonify{
			"message": err,
		})
	} else {
		nowUser := MatchingModel{
			Sid:    s.ID(),
			UserId: user.Id,
		}
		s.Emit("join", jsonify{
			"message": "进入房间成功",
		})
		if len(matching) != 0 {
			match(nowUser, matching[0])
		} else {
			exist := false
			for _, v := range matching {
				if v.UserId == nowUser.UserId {
					exist = true
				}
			}
			if exist == false {
				matching = append(matching, nowUser)
			}
		}
	}
}

func Exit(s socketio.Conn) {
	j := 0
	exist := false
	for _, val := range matching {
		if val.Sid != s.ID() {
			exist = true
			matching[j] = val
			j++
		}
	}
	matching = matching[:j]
	if exist {
		s.Emit("exit", jsonify{
			"message": "退出房间成功",
		})
	}
}

// 匹配
func match(s1, s2 MatchingModel) {

	if s1.UserId == s2.UserId {
		return
	} else {
		count, err := db.Count(&QuestionModel{})
		strInt64 := strconv.FormatInt(count, 10)
		count32, _ := strconv.Atoi(strInt64)
		question := QuestionModel{}
		var has bool
		for !has {
			rand.Seed(time.Now().UnixNano())
			num := rand.Intn(count32) + 1 //[1,count]
			question.Id = num
			has, err = db.Get(&question)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
		picture := PictureModel{
			UserId1:  s1.UserId,
			UserId2:  s2.UserId,
			Question: question.Name,
		}
		user1 := UserModel{Id: picture.UserId1}
		user2 := UserModel{Id: picture.UserId2}
		_, err = db.Get(&user1)
		_, err = db.Get(&user2)
		if err != nil {
			log.Println(err.Error())
			return
		}
		db.Insert(picture)
		matching = matching[1:]
		broadcast.Send(s1.Sid, "match", jsonify{
			"message": "匹配成功",
			"data": jsonify{
				"another_user_name": user2.Name,
				"question":          question.Id,
				"position":          "top",
			},
			// 其他再加
		})
		broadcast.Send(s2.Sid, "match", jsonify{
			"message": "匹配成功",
			"data": jsonify{
				"another_user_name": user1.Name,
				"question":          question.Id,
				"position":          "bottom",
			},
			// 其他再加
		})
	}
}
