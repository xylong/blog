package middleware

import (
	"blog/internal"
	"blog/internal/util"
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
		authorization := c.Request.Header.Get(authKey)
		list := strings.Split(authorization, sep)
		if len(list) != length || list[0] != tokenType {
			c.Status(http.StatusUnauthorized)
			c.Set("code", internal.TokenMalformed)
			c.Abort()
			return
		}

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
