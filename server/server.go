package server

import (
	"fmt"
	"net/http"
	"tomaxut/internal/app/api"
	"tomaxut/server/middleware"
	"tomaxut/server/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitRouter(router *gin.Engine) *gin.Engine {

	err := router.SetTrustedProxies([]string{"192.168.20.55"})
	if err != nil {
		return nil
	}
	router.GET("/", func(c *gin.Context) {
		// If the client is 192.168.1.2, use the X-Forwarded-For
		// header to deduce the original client IP from the trust -
		// worthy parts of that header.
		// Otherwise, simply return the direct client IP
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
		logrus.Warning("ClientIP: %s\n", c.ClientIP())
		response.Success(c, &response.SuccessResponse{
			Code: http.StatusOK,
		})
	})

	v1 := router.Group("/api/v1")
	{
		rabbit := new(api.Rabbit)
		v1.GET("/rabbit", rabbit.SendMessage)
		auth := new(api.Auth)
		v1.POST("/login", auth.Login)
		v1.POST("/register", auth.Register)

		v1.Use(middleware.JWT("user"))
		v1.Use(middleware.Auth())

		v1.PATCH("/change-password", auth.ChangePassword)
		v1.PATCH("/frozen-account", auth.FrozenAccount)

		user := new(api.User)
		v1.GET("/user", user.Index)
		v1.POST("/user", user.Store)
		v1.PATCH("/user/:id", user.Update)
		v1.DELETE("/user/:id", user.Destroy)

		post := new(api.Post)
		v1.GET("/post", post.Index)
		v1.GET("/post/:id", post.Show)
		v1.POST("/post", post.Store)
	}
	return router

}
