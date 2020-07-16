package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		logrus.Infof("| %3d | %13v | %15s | %s  %s |",
			c.Writer.Status(),
			end.Sub(start),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}
