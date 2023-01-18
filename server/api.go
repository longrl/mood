package main

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/controller"
	"github.com/longrl/mood/middleware"
)

func RegisterRoute(route *gin.Engine) {
	blog := route.Group("blog")
	{
		blog.POST("/authority", controller.Authority)
		blog.GET("/captcha", middleware.Auth(), controller.ShowCaptcha)
	}
}
