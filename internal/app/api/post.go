package api

import (
	"tomaxut/internal/app/requests"
	"tomaxut/internal/app/services"
	"tomaxut/pkg/utils/validator"
	"tomaxut/server/response"

	"github.com/gin-gonic/gin"
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

func (*Post) Show(ctx *gin.Context) {
	validated := validator.Validate(ctx, &requests.DetailRequest{})

	if len(validated) > 0 {
		response.UnprocessableEntity(ctx, validated)
		return
	}
	result, err := services.NewPost().Find(ctx)
	if err != nil {
		response.Fail(ctx, err)
		return
	}
	response.Success(ctx, result)
}
