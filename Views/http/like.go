package http

import "github.com/gin-gonic/gin"

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

}

func Rank(c *gin.Context) (*map[string]interface{}, int, error) {

}
