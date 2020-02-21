package api

import (
	"Go/pkg/tgin/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

/**
自定义中间件
 */
func CustomMiddleware(router *gin.Engine) {
	router.Use(middleware.Logger())
	router.GET("/custom-middleware", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		// 它将打印： "12345"
		log.Println(example)
	})
}
