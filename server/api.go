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
		blog.GET("/captcha", middleware.Auth(), middleware.LimitPerRoute("50-H"), controller.ShowCaptcha)

		blog.GET("/detail/:id", controller.Detail)
		blog.GET("/list", controller.List)
		blog.GET("/archive", controller.Archive)
		blog.GET("/delete/:id", controller.Delete)
		blog.POST("/add", middleware.Auth(), controller.Add)
		blog.POST("/detail", middleware.Auth(), controller.UpdateDetail)
		blog.POST("/top/:id", middleware.Auth(), controller.Top)
	}
}
