package api

import "github.com/gin-gonic/gin"

/**
默认的 Gin
	已经连接了 Logger 和 Recovery 中间件
 */
func DefaultRouter() (router *gin.Engine) {
	return gin.Default()
}

/**
空白的没有任何中间件的 Gin
 */
func SimpleRouter() (router *gin.Engine) {
	return gin.New()
}




