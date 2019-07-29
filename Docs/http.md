<h1 align="center">HTTP接口文档</h1>

## 登录
- url: /api/auth
- method: POST
- args: 

| args | nullable | type | remark   |
|:------:|:------:|:------:|:------:|
|    student_id    |    false    |    string   |    学号    |
|    name    |    false    |    string   |    姓名      |
- return:

      {
          "message": "登录成功",
          "status": 1,
          "data": {
              "token": token
          }
      }
      
      //有项目为空
      {
          "message": "学号和姓名不能为空"，
          "status": 0
      }, 400
      
      //发生其他错误时
      {
          "message": 错误讯息，
          "status": 0
      }, 400
