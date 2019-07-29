package http

import (
	. "PrintHalf/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PictureView(g *gin.RouterGroup) {
	// 获取分页图画
	g.GET("/all", func(c *gin.Context) {
		if result, status, err := GetPictures(c); err != nil {
			c.AbortWithStatusJSON(status, result)
		} else {
			c.JSON(status, result)
		}
	})
	// 获取特定图画
	g.GET("", func(c *gin.Context) {
		if result, status, err := GetPicture(c); err != nil {
			c.AbortWithStatusJSON(status, result)
		} else {
			c.JSON(status, result)
		}
	})
}

func GetPicture(c *gin.Context) (*map[string]interface{}, int, error) {
	id, _ := strconv.Atoi(c.Query("page"))
	picture := PictureModel{
		Id: id,
	}
	has, err := db.Get(&picture)
	if err != nil {
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusBadRequest, err
	}
	if !has {
		return &map[string]interface{}{
			"message": "该图不存在",
			"status":  0,
		}, http.StatusBadRequest, err
	}
	user1, user2 := UserModel{
		Id: picture.UserId1,
	}, UserModel{
		Id: picture.UserId2,
	}
	db.Get(user1)
	db.Get(user2)
	return &map[string]interface{}{
		"message": "成功",
		"status":  1,
		"data": map[string]interface{}{
			"name1":            user1.Name,
			"name2":            user2.Name,
			"top_file_name":    picture.TopFileName,
			"bottom_file_name": picture.BottomFileName,
			//"like_num": picture.LikeNum,
		},
	}, http.StatusOK, nil
}

func GetPictures(c *gin.Context) (*map[string]interface{}, int, error) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pictures := make([]PictureModel, 0)
	db.Limit(12, 12*(page-1)).Find(&pictures)
	if len(pictures) == 0 {
		return &map[string]interface{}{
			"message": "该页不存在",
			"status":  0,
		}, http.StatusBadRequest, nil
	}
	total, _ := db.Count(PictureModel{})
	total_page := total / 12
	var pictures_data [](map[string]interface{})
	for _, picture := range pictures {
		//var picture_data map[string]interface{}
		user1 := UserModel{Id: picture.UserId1}
		user2 := UserModel{Id: picture.UserId2}
		db.Get(user1)
		db.Get(user2)
		picture_data := map[string]interface{}{
			"name1":            user1.Name,
			"name2":            user2.Name,
			"top_file_name":    picture.TopFileName,
			"bottom_file_name": picture.BottomFileName,
		}
		pictures_data = append(pictures_data, picture_data)
	}
	return &map[string]interface{}{
		"message": "成功",
		"status":  1,
		"data": map[string]interface{}{
			"pictures_data": pictures_data,
		},
		"total_page": total_page,
	}, http.StatusOK, nil
}
