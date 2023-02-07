package services

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"tomaxut/pkg/utils/ginparam"
	"tomaxut/pkg/utils/paginate"
	"tomaxut/server/response"
	"tomaxut/store/models"
	"tomaxut/store/repositories"
)

type User struct {
	repo *repositories.User
}

func (u *User) Query(c *gin.Context) (*response.SuccessResponse, error) {
	paginateParam := paginate.RequestParam(c)
	items, total, err := u.repo.Paginate(paginateParam)

	return &response.SuccessResponse{
		Data: items,
		Meta: &paginate.Meta{
			CurrentPage: paginateParam.Page,
			PerPage:     paginateParam.PageSize,
			Total:       total,
		},
	}, err
}

func (u *User) Store(c *gin.Context) error {
	var m models.User
	if err := c.ShouldBind(&m); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.Password = string(hash)

	return u.repo.Create(&m)
}

func (u *User) Update(c *gin.Context) error {
	data := ginparam.PostKeysToMap(c, []string{
		"name", "email", "sex", "status",
	})

	return u.repo.Update(c.Param("id"), data)
}

func (u *User) Delete(c *gin.Context) error {
	return u.repo.Delete(c.Param("id"))
}

func NewUser() *User {
	return &User{
		repo: &repositories.User{},
	}
}
