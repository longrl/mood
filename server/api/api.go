package api

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/controller"
	"github.com/longrl/mood/middleware"
)

func RegisterRoute(route *gin.Engine) {
	route.Use(middleware.Cors())
	route.MaxMultipartMemory = 8 << 20 // 8 MiB
	blog := route.Group("blog")
	{
		blog.POST("/authority", controller.Authority)
		blog.GET("/captcha", middleware.LimitPerRoute("600-H"), controller.ShowCaptcha)

		blog.GET("/detail/:id", middleware.LimitPerRoute("10000-H"), controller.Detail)
		blog.GET("/list", middleware.LimitPerRoute("10000-H"), controller.List)
		blog.GET("/archive", middleware.LimitPerRoute("10000-H"), controller.Archive)
		blog.GET("/delete/:id", middleware.Auth(), controller.Delete)
		blog.GET("/top", middleware.LimitPerRoute("10000-H"), controller.GetTop)
		blog.POST("/add", middleware.Auth(), controller.Add)
		blog.POST("/detail", middleware.Auth(), controller.UpdateDetail)
		blog.POST("/top/:id", middleware.Auth(), controller.Top)
		blog.POST("/upload", controller.Upload)
	}
}
