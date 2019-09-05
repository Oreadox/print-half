package http

import (
	"PrintHalf/Config"
	. "PrintHalf/Models"
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
	user, _ := c.Get("user")
	userId := user.(UserModel).Id
	var picture PictureModel
	has, err := db.Where("user_id1 = ?", userId).Or("user_id2 = ?", userId).OrderBy("id DESC").Get(&picture)
	if err != nil {
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusInternalServerError, err
	} else if !has {
		return &map[string]interface{}{
			"message": "房间不存在",
			"status":  0,
		}, http.StatusOK, nil
	}
	file, _, err := c.Request.FormFile("image") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	format := c.Request.FormValue("format")
	if err != nil {
		log.Println(err)
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusBadRequest, err
	}
	filename := utils.GetRandomString(16) + "." + format
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
	if picture.UserId1 == userId {
		picture.TopFileName = filename
		_, err = db.Id(picture.Id).Cols("top_file_name").Update(picture)
		if err != nil {
			return &map[string]interface{}{
				"message": err.Error(),
				"status":  0,
			}, http.StatusInternalServerError, err
		}
	} else if picture.UserId2 == userId {
		picture.BottomFileName = filename
		_, err = db.Id(picture.Id).Cols("bottom_file_name").Update(picture)
		if err != nil {
			return &map[string]interface{}{
				"message": err.Error(),
				"status":  0,
			}, http.StatusInternalServerError, err
		}
	}
	os.Remove("./static/uploadfile/" + filename)
	return &map[string]interface{}{
		"message": "成功",
		"status":  1,
		"data": &map[string]interface{}{
			"picture_id": picture.Id,
			"filename":   filename,
		},
	}, http.StatusOK, nil
}

func SaveFile(file multipart.File, filename string) (*os.File, error) {
	os.MkdirAll("static/uploadfile/", os.ModePerm)
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
