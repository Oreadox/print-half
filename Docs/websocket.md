<h1 align="center">WebSocket接口文档</h1>

## 进入游戏
- input

      [
          "join",
          {
              "token": token
          }
      ]
      
- return

      [
          "join",
          {
              "message": "进入房间成功"   //匹配成功在下面
          }
      ]
      
      //token无效时
      [
          "join",
          {
              "message": "token无效"   //匹配成功在下面
          }
      ]
      
## 匹配成功
* input

  无（由服务器主动推送）
    
- return

      [
          "match",
          {
                    "message": "匹配成功",
                    "data":{
                        "another_user_name": 另一个用户的名字，
                        "question": 题目id
                        "position": "top" or "bottom" (位置)
                    }
          }
      ]
      //两个用户都会发
      
## 退出匹配
- input

      [
          "exit"
      ]
      
- return      
      
      [
          "exit",
          {
              "message": "退出房间成功"
          }
      ]
      
      
## 进入下一回合
- input

      [
          "next",
          {
              "token": token        //token需为后台用户的
          }
      ]
      
      
- return
        
        
      // 向后台用户
      [
          "next",
          {
              "message": "成功"
          }
      ]
      
      
      // 所有已连接的用户
      [
          "next_round",
          {
              "message": "下一回合开始"
              "data": {
                  "now_round": 当前回合
              }
          }
      ]             
 