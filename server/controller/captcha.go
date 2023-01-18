package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/config/captcha"
	"github.com/longrl/mood/response"
)

func ShowCaptcha(c *gin.Context) {
	id, b64, _ := captcha.NewCaptcha().GenerateCaptcha()
	response.JSON(c, gin.H{
		"id":  id,
		"b64": b64,
	})
}
