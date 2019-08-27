package http

import (
	"PrintHalf/Config"
	utils "PrintHalf/Utils"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func RoomView(g *gin.RouterGroup) {
	g.POST("/upload", func(c *gin.Context) {
		if result, status, err := Upload(c); err != nil {
			c.AbortWithStatusJSON(status, result)
		} else {
			c.JSON(status, result)
		}
	})
}

func Upload(c *gin.Context) (*map[string]interface{}, int, error) {
	file, _, err := c.Request.FormFile("image") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		log.Println(err)
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusBadRequest, err
	}
	filename := utils.GetRandomString(16)
	out, err := SaveFile(file, filename)
	if err != nil {
		log.Println(err)
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusInternalServerError, err
	}
	err = UploadFile(filename, out)
	if err != nil {
		log.Println(err)
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusInternalServerError, err
	}
	os.Remove("./static/uploadfile/" + filename)
	return &map[string]interface{}{
		"message": "成功",
		"status":  1,
	}, http.StatusOK, nil
}

func SaveFile(file multipart.File, filename string) (*os.File, error) {
	os.MkdirAll("/static/uploadfile/", 0755)
	out, err := os.Create("static/uploadfile/" + filename)
	if err != nil {
		return nil, err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	out, err = os.Open("./static/uploadfile/" + filename)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func UploadFile(filename string, fd *os.File) error {
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	bucket, err := client.Bucket(config.BucketName)
	if err != nil {
		return err
	}
	err = bucket.PutObject("picture/"+filename, fd)
	if err != nil {
		return err
	}
	return nil
}
