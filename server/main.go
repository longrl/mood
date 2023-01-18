package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	// 注册路由
	RegisterRoute(router)

	router.Run(":8080")
}
