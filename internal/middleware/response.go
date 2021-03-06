package middleware

import (
	"blog/internal"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
)

func ResponseHandler(c *gin.Context) {
	c.Next()
	if c.Writer.Status() == http.StatusNotFound && c.Writer.Size() <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "找不到资源",
		})
		return
	}
	if c.Writer.Status() == http.StatusOK && c.Writer.Size() <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		return
	}
}

func Recovery(c *gin.Context) {
	httpCode := http.StatusOK

	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 65536)
			buf = buf[:runtime.Stack(buf, false)]
			e, ok := err.(*internal.Error)
			if ok {
				// 翻译
				errs, ok := e.Err.(validator.ValidationErrors)
				if ok {
					// 设置http为400
					httpCode = http.StatusBadRequest
					trans, _ := c.Value(TranslatorKey).(ut.Translator)
					for _, item := range errs {
						e.Msg = item.Translate(trans)
						break
					}
				}

				if e.Code >= http.StatusInternalServerError {
					logrus.Errorf("%s\n%s", err, buf)
				}

				c.AbortWithStatusJSON(httpCode, gin.H{
					"code": e.Code,
					"msg":  e.Msg,
					"data": gin.H{},
				})

				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "服务出错，请稍后尝试",
			})
		}
	}()

	c.Next()
}
