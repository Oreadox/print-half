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
                        "question": 题目
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
 