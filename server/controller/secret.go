package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/config"
	"github.com/longrl/mood/request"
	"github.com/longrl/mood/response"
	"github.com/longrl/mood/service"
	"time"
)

func Authority(c *gin.Context) {
	req := request.SecretRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	secretService := service.SecretService{}
	ok, s := secretService.Check(req.Key)
	// 将secret数据存入session中或者redis中
	if ok {
		// 设置授权时间为12h
		config.Redis.Set(c.RemoteIP(), s.Role, 12*time.Hour)
		response.Success(c)
	} else {
		response.Abort403(c)
	}
}