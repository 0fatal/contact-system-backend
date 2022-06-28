package middleware

import (
	"backend/response"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/api/login" {
			c.Next()
			return
		}

		if _, err := c.Cookie("user"); err != nil {
			response.Fail().Msg("需要登录").Send(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
