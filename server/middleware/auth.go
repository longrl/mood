package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/config"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := config.Redis.Get(c.ClientIP())
		if role != "2" {
			c.AbortWithStatusJSON(200, gin.H{
				"code":    1001,
				"message": "权限不足",
			})
		}
		c.Next()
	}
}
