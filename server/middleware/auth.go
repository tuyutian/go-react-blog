package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tomaxut/pkg/auth"
	"tomaxut/server/response"
	"tomaxut/store/database"
	"tomaxut/store/models"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := database.DB.Where("id=? and status=0", auth.New().JwtUserId(c)).First(&user).Error; err != nil {
			response.Fail(c, &response.FailResponse{
				Code:    http.StatusUnauthorized,
				Message: "用户不存在",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
