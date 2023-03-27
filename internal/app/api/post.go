package api

import (
	"github.com/gin-gonic/gin"
	"tomaxut/internal/app/requests"
	"tomaxut/internal/app/services"
	"tomaxut/pkg/utils/validator"
	"tomaxut/server/response"
)

type Post struct {
}

func (*Post) Index(ctx *gin.Context) {
	result, err := services.NewPost().Query(ctx)
	if err != nil {
		response.Fail(ctx, err)
	}
	response.Success(ctx, result)
}

func (*Post) Store(ctx *gin.Context) {
	result := validator.Validate(ctx, &requests.PostRequest{})

	if len(result) > 0 {
		response.UnprocessableEntity(ctx, result)
		return
	}

	if err := services.NewUser().Store(ctx); err != nil {
		response.Fail(ctx, err)
		return
	}

	response.Created(ctx)
}
