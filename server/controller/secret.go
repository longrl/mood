package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/config"
	"github.com/longrl/mood/config/captcha"
	"github.com/longrl/mood/request"
	"github.com/longrl/mood/response"
	"github.com/longrl/mood/service"
	"time"
)

func Authority(c *gin.Context) {
	req := request.SecretRequest{}
	err := c.ShouldBind(&req)
	if err != nil || req.Id == "" {
		response.Error(c, errors.New("mood: 传入参数有问题"))
		return
	}

	ok := captcha.NewCaptcha().VerifyCaptcha(req.Id, req.Answer)
	if !ok {
		response.JSON(c, gin.H{
			"code":    999,
			"message": "验证失败",
		})
		return
	}
	secretService := service.SecretService{}
	ok, s := secretService.Check(req.Key)
	// 将secret数据存入session中或者redis中
	if ok {
		// 设置授权时间为12h
		config.Redis.Set(c.ClientIP(), s.Role, 12*time.Hour)
		response.Success(c)
	} else {
		response.Abort403(c)
	}
}
