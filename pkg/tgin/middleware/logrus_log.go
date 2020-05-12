package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func LogrusLog(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}

		log.WithFields(logrus.Fields{
			"host_name":   hostName,
			"status_code": c.Writer.Status(),
			"latency":     latency,
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"user_agent":  c.Request.UserAgent(),
			"referer":     c.Request.Referer(),
			"data_length": c.Writer.Size(),
			"path":        path,
			//"params":      GetParams(c),
		}).Info("request api")
	}
}
