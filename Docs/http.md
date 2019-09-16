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
      
      //学号与姓名不对应
      {
          "message": "学号与姓名不对应"，
          "status": 0
      }, 200
      
      //有项目为空
      {
          "message": "学号和姓名不能为空"，
          "status": 0
      }, 200
      
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
|    num    |    false    |    int   |    点赞数目(1或-1)    |          

- return:

      {
          "message": "成功",
          "status": 1
      }, 200
      
      //参数错误(非有效值或取消点赞后该图片赞数将为负)
      {
          "message": "无效的参数",
          "status": 0
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
      
## 获取图片

- url: https://draw2019.oss-cn-shanghai.aliyuncs.com/picture/pic_name
- url(指定大小): https://draw2019.oss-cn-shanghai.aliyuncs.com/picture/pic_name?x-oss-process=image/resize,m_fill,h_高度,w_宽度,limit_0
- method: GET

## 上传图片
先将图片上传至服务器，服务器会把图片上传到oss上

- url: /api/room/upload
- method: POST
- request_header: Authorization: token
- args:      

| args | nullable | type | remark   |
|:------:|:------:|:------:|:------:|
|    img    |    false    |    str   |    图片base64    |
|    format    |    false    |    str   |    图片格式    |

      
- return:

      {
          "message": "成功",
          "status": 1,
          "data":{
              "picture_id": 图片id,
              "filename": 当前用户所画的图片的文件名
          }
      }, 200

      //找不到用户所在的房间
      {
          "message": "房间不存在"，
          "status": 0
      }, 404      

      //发生其他错误
      {
          "message": 错误讯息，
          "status": 0
      }, 500
