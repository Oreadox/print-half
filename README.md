# 两人画

## 初始化

- 包的安装

      go get github.com/googollee/go-socket.io
      go get github.com/dgrijalva/jwt-go
      go get github.com/gin-gonic/gin
      go get github.com/go-xorm/xorm
      go get github.com/go-sql-driver/mysql
      go get github.com/aliyun/aliyun-oss-go-sdk/oss
      go get github.com/gin-contrib/cors
      
- 环境变量设置      
    
      SecretKey: 加密密钥
      DbUri: 数据库URL
      ExpiresTime: Token过期时间
      //OSS设置:
      AccessKeyId
      AccessKeySecret
      Endpoint
      BucketName
      // 上述设置修改Config/config.go文件也可

## 接口文档
- [HTTP](https://github.com/yangchen29/print-half/blob/master/Docs/http.md)
- [WebSocket](https://github.com/yangchen29/print-half/blob/master/Docs/WebSocket.md)      

