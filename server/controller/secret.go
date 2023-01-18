package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/service"
)

type Secret struct {
	Key string `json:"key"`
}

func Authority(c *gin.Context) {
	req := Secret{}
	c.ShouldBind(&req)
	service := service.SecretService{}
	ok, _ := service.Check(req.Key)
	// 将secret数据存入session中
	if ok {
		c.JSON(200, gin.H{
			"msg": "successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "error",
		})
	}
}
