package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"toDoListDemo/utils"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		var code int

		token := c.GetHeader("Authorization") //从请求头中拿到这个字段的值
		if token == "" {
			code = 400
		} else {
			claims, err := utils.ParseToken(token) //调用utils包中的ParseToken方法，解析token
			if err != nil {
				code = 403
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401
			}
		}

		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
