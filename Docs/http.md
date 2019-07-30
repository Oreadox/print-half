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
      }, 200
      
      //有项目为空
      {
          "message": "学号和姓名不能为空"，
          "status": 0
      }, 400
      
      //发生其他错误
      {
          "message": 错误讯息，
          "status": 0
      }, 500

## 分页的所有图片信息
- url: /api/picture/all
- method: GET
- request_header: Authorization: token
- args:

| args | nullable | type | remark   |
|:------:|:------:|:------:|:------:|
|    page    |    true    |    int   |    页号(默认为1）    |

- return:

      {
          "message": "成功",
          "status": 1,
          "data": {
              "pictures_data": [
                  {
                      "id": 图片id,
                      "name1": 用户1姓名,
                      "name2": 用户2姓名，
                      "top_file_name": 顶部图片文件名,
                      "bottom_file_name": 底部图片文件名  
                  },
                  ...
              ]
          },
          "total_page": 总页数
      }, 200
      // 一页暂定12张
      
      //该页不存在
      {
          "message": "该页不存在"，
          "status": 0
      }, 404
      
      //发生其他错误
      {
          "message": 错误讯息，
          "status": 0
      }, 500
      
## 单个图片信息
- url: /api/picture
- method: GET
- request_header: Authorization: token
- args:

| args | nullable | type | remark   |
|:------:|:------:|:------:|:------:|
|    id    |    false    |    int   |    图片id    |

- return:

      {
          "message": "成功",
          "status": 1,
          "data": {
              "name1": 用户1姓名,
              "name2": 用户2姓名，
              "top_file_name": 顶部图片文件名,
              "bottom_file_name": 底部图片文件名  
          }
      }, 200
      
      //该图不存在
      {
          "message": "该图不存在"，
          "status": 0
      }, 404
      
      //发生其他错误
      {
          "message": 错误讯息，
          "status": 0
      }, 500

## 图片点赞
- url: /api/like
- method: POST
- request_header: Authorization: token
- args:      

| args | nullable | type | remark   |
|:------:|:------:|:------:|:------:|
|    id    |    false    |    int   |    图片id    |      

- return:

      {
          "message": "成功",
          "status": 1
      }, 200
      
      //该图不存在
      {
          "message": "该图不存在"，
          "status": 0
      }, 404
      
      //发生其他错误
      {
          "message": 错误讯息，
          "status": 0
      }, 500

## 图片赞数排名

- url: /api/like/rank
- method: GET
- request_header: Authorization: token
* args:      
    
  none
- return:

      {
          "message": "成功",
          "status": 1,
          "data": {
              "pictures_data": [
                  {
                      "id": 图片id,
                      "name1": 用户1姓名,
                      "name2": 用户2姓名，
                      "top_file_name": 顶部图片文件名,
                      "bottom_file_name": 底部图片文件名  
                  },
                  ...
              ]
          },
          "total_page": 总页数
      }, 200
      // 列表已排好顺序
      
      //发生其他错误
      {
          "message": 错误讯息，
          "status": 0
      }, 500
      