package services

import (
	"strconv"
	"tomaxut/pkg/auth"
	"tomaxut/pkg/utils/paginate"
	"tomaxut/server/response"
	"tomaxut/store/models"
	"tomaxut/store/repositories"

	"github.com/gin-gonic/gin"
)

type Post struct {
	repo *repositories.Post
}

func (p *Post) Query(c *gin.Context) (*response.SuccessResponse, error) {
	paginateParam := paginate.RequestParam(c)
	items, total, err := p.repo.Paginate(paginateParam)

	return &response.SuccessResponse{
		Data: items,
		Meta: &paginate.Meta{
			CurrentPage: paginateParam.Page,
			PerPage:     paginateParam.PageSize,
			Total:       total,
		},
	}, err
}

func (p *Post) Store(c *gin.Context) error {
	var m models.Post
	if err := c.ShouldBind(&m); err != nil {
		return err
	}
	m.Content = c.PostForm("content")
	m.Subtitle = c.PostForm("subtitle")
	m.Title = c.PostForm("title")
	m.UserId = auth.New().JwtUserId(c)
	return p.repo.Create(&m)
}

func (p *Post) Find(c *gin.Context) (*response.SuccessResponse, error) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		return &response.SuccessResponse{}, err
	}

	result, err := p.repo.Find(id)
	return &response.SuccessResponse{
		Data: result,
	}, err
}

func NewPost() *Post {
	return &Post{
		repo: &repositories.Post{},
	}
}
