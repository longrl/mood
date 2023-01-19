package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/longrl/mood/config"
	"github.com/longrl/mood/response"
	"net/http"
)

// LimitPerRoute 限流中间件，用在单独的路由中
func LimitPerRoute(limit string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 针对单个路由，增加访问次数
		c.Set("limiter-once", false)

		// 针对 IP + 路由进行限流
		key := config.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, limit string) bool {

	// 获取超额的情况
	rate, err := config.CheckRate(c, key, limit)
	if err != nil {
		response.Abort500(c)
		return false
	}

	// ---- 设置标头信息-----
	// X-RateLimit-Limit :10000 最大访问次数
	// X-RateLimit-Remaining :9993 剩余的访问次数
	// X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-RateLimit-Limit
	c.Header("X-RateLimit-Limit", "10000")
	c.Header("X-RateLimit-Remaining", "9993")
	c.Header("X-RateLimit-Reset", "1513784506")

	// 超额
	if rate.Reached {
		// 提示用户超额了
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "接口请求太频繁",
		})
		return false
	}

	return true
}
