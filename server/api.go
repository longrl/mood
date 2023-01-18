package main

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/controller"
)

func RegisterRoute(route *gin.Engine) {
	blog := route.Group("blog")
	{
		blog.POST("/authority", controller.Authority)
	}
}
