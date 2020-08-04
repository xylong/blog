package middleware

import (
	"blog/internal"
	"blog/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

const (
	authKey   = "Authorization"
	sep       = " "
	tokenType = "Bearer"
	length    = 2
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		authorization := c.Request.Header.Get(authKey)
		if authorization == "" {
			c.Status(http.StatusUnauthorized)
			c.Set("code", internal.UNAUTHORIZED)
			internal.PanicError(http.StatusUnauthorized, errors.New(internal.GetMsg(internal.UNAUTHORIZED)))
			return
		}
		// 判断token格式
		list := strings.Split(authorization, sep)
		if len(list) != length || list[0] != tokenType {
			c.Status(http.StatusUnauthorized)
			c.Set("code", internal.TokenMalformed)
			internal.PanicError(http.StatusUnauthorized, errors.New(internal.GetMsg(internal.TokenMalformed)))
		}
		// 解析
		claims, err := util.NewJWT().Parse(list[1])
		if err != nil {
			internal.PanicError(http.StatusUnauthorized, err)
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.Set("code", internal.TokenExpired)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
