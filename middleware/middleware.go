package middleware

import (
	"context"
	"go-pangu/controller"
	"go-pangu/db"
	"go-pangu/jwt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//中间件 验证token scp传入用户类型
func Auth(scp string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bear := c.Request.Header.Get("Authorization")
		token := strings.Replace(bear, "Bearer ", "", 1)
		sub, scope, err := jwt.Decoder(token)
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, err.Error())
		} else {
			if scope != scp {
				controller.StatusError(c, http.StatusUnauthorized, "unauthorized", "invalid scope")
				c.Abort()
			}
			c.Set("sub", sub)
			c.Set("scp", scope)
			c.Next()
		}
	}
}

func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmieOutContext, _ := context.WithTimeout(context.Background(), time.Second)
		ctx := context.WithValue(c.Request.Context(), "DB", db.DB.WithContext(tmieOutContext))
		c.Request.WithContext(ctx)
		c.Next()
	}
}
