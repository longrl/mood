package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/config"
	"github.com/longrl/mood/response"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := config.Redis.Get(c.ClientIP())
		if role != "2" {
			response.Abort403(c)
		}
		c.Next()
	}
}
