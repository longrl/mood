package main

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/config"
	"github.com/longrl/mood/model/secret"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// 初始化DB
	config.SetupDB()
	config.DB.AutoMigrate(&secret.Secret{})
	// 注册路由
	RegisterRoute(router)

	router.Run(":8080")
}
