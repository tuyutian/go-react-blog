package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tomaxut/pkg/auth"
	"tomaxut/server/response"
)

func JWT(guard string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = token[7:]
		var message string
		if token == "" {
			message = "token不能为空"
		} else {
			claims, err := new(auth.JwtToken).ParseToken(token)
			if err != nil {
				message = err.Error()
			} else if guard != claims.Guard {
				message = "Guard 有误"
			}
		}

		if len(message) > 0 {
			response.Fail(c, &response.FailResponse{Code: http.StatusUnauthorized, Message: message})
			c.Abort()
			return
		}

		c.Next()
	}
}
