package http

import (
	. "PrintHalf/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	var likeInfo LikeModel
	err := c.ShouldBindJSON(&likeInfo)
	if err != nil {
		return &map[string]interface{}{
			"message": err.Error(),
			"status":  0,
		}, http.StatusOK, err
	}
	if !(likeInfo.Num == 1 || likeInfo.Num == -1) {
		return &map[string]interface{}{
			"message": "无效的参数",
			"status":  0,
		}, http.StatusOK, nil
	}
	picture := PictureModel{
		Id: likeInfo.Id,
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
	if picture.LikeNum <= 0 && likeInfo.Num < 0 {
		return &map[string]interface{}{
			"message": "无效的参数",
			"status":  0,
		}, http.StatusOK, nil
	}
	picture.LikeNum += likeInfo.Num
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
	num, _ := strconv.Atoi(c.DefaultQuery("num", "3"))
	pictures := make([]PictureModel, 0)
	roundSetting := SettingModel{
		Desc: "NowRound",
	}
	db.Get(&roundSetting)
	err := db.Where("round = ?", roundSetting.Value).OrderBy("like_num desc").Limit(num, 0).Find(&pictures)
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
			"like_num":         picture.LikeNum,
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
