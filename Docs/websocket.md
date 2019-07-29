<h1 align="center">HTTP接口文档</h1>

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
      
      