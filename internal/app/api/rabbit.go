package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tomaxut/server/response"
	"tomaxut/store/database"
)

type Rabbit struct {
}

func (r *Rabbit) SendMessage(ctx *gin.Context) {
	rabbitmq := database.NewRabbitMQSimple("" +
		"tomaxut")
	rabbitmq.PublishSimple("Hello tomaxut ,you are success!")
	fmt.Println("发送成功！")

	response.Success(ctx, &response.SuccessResponse{
		Code:    200,
		Status:  "200",
		Message: "success",
	})
}
