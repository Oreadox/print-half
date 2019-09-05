package http

import (
	. "PrintHalf/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LikeView(g *gin.RouterGroup) {
	// 获取当前总排名
	g.GET("/rank", func(c *gin.Context) {
		if result, status, err := Rank(c); err != nil {
			c.AbortWithStatusJSON(status, result)
		} else {
			c.JSON(status, result)
		}
	})

	g.POST("", func(c *gin.Context) {
		if result, status, err := Like(c); err != nil {
			c.AbortWithStatusJSON(status, result)
		} else {
			c.JSON(status, result)
		}
	})
}

func Like(c *gin.Context) (*map[string]interface{}, int, error) {
	var picture_id struct {
		id int
	}
	err := c.ShouldBindJSON(&picture_id)
	if err != nil {
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusOK, err
	}
	picture := PictureModel{
		Id: picture_id.id,
	}
	has, err := db.Get(&picture)
	if !has {
		return &map[string]interface{}{
			"message": "图片不存在",
			"status":  0,
		}, http.StatusNotFound, err
	} else if err != nil {
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusInternalServerError, err
	}
	picture.LikeNum += 1
	_, err = db.Id(picture.Id).Cols("like_num").Update(picture)
	if err != nil {
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusInternalServerError, err
	}
	return &map[string]interface{}{
		"message": "成功",
		"status":  1,
	}, http.StatusOK, nil
}

func Rank(c *gin.Context) (*map[string]interface{}, int, error) {
	pictures := make([]PictureModel, 0)
	err := db.OrderBy("like_num desc").Limit(3, 0).Find(&pictures)
	if err != nil {
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusInternalServerError, err
	}
	var pictures_data [](map[string]interface{})
	for _, picture := range pictures {
		//var picture_data map[string]interface{}
		user1 := UserModel{Id: picture.UserId1}
		user2 := UserModel{Id: picture.UserId2}
		db.Get(&user1)
		db.Get(&user2)
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
	}, http.StatusOK, nil
}
