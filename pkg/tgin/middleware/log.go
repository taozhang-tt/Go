package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

/**
自定义log中间件
 */
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置简单的变量
		c.Set("example", "12345")

		// 在请求之前

		c.Next()

		// 在请求之后
		latency := time.Since(t)
		log.Print(latency)

		// 记录我们的访问状态
		status := c.Writer.Status()
		log.Println(status)
	}
}