package websocket

import (
	. "PrintHalf/Models"
	utils "PrintHalf/Utils"
	"fmt"
	"github.com/googollee/go-socket.io"
	jsoniter "github.com/json-iterator/go"
)

func Next(s socketio.Conn, data string) {
	var tokenData struct {
		Token string
	}
	playerBroadcast.Join(s.ID(), s)
	jsoniter.Unmarshal([]byte(data), &tokenData)
	token := tokenData.Token
	fmt.Println("token:" + token)
	user, err := utils.VerifyAuthToken(token)
	if err != "" {
		s.Emit("next", jsonify{
			"message": err,
		})
	}
	if user.Name != "test" {
		s.Emit("next", jsonify{
			"message": "无效后台用户",
		})
	} else {
		roundSetting := SettingModel{
			Desc: "NowRound",
		}
		db.Get(&roundSetting)
		roundSetting.Value += 1
		_, err := db.Id(roundSetting.Id).Cols("value").Update(roundSetting)
		if err != nil {
			s.Emit("next", jsonify{
				"message": err.Error(),
			})
		} else {
			s.Emit("next", jsonify{
				"message": "成功",
			})
			userBroadcast.Send("allPlayer", "next_round", jsonify{
				"message": "下一回合开始",
				"data": jsonify{
					"now_round": roundSetting.Value,
				},
			})
		}

	}

}
