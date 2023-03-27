package repositories

import (
	"tomaxut/pkg/utils/paginate"
	"tomaxut/store/database"
	"tomaxut/store/models"
)

type Post struct {
	Helper
}

func (p *Post) Paginate(paginateParam *paginate.Param) ([]models.Post, int64, error) {
	var items []models.Post
	total := p.Count()

	result := database.DB.Scopes(paginate.ORMScope(paginateParam)).Find(&items)

	return items, total, result.Error
}

func (*Post) Get() ([]models.Post, error) {
	var items []models.Post
	result := database.DB.Model(models.Post{}).Find(items)
	return items, result.Error
}

func (*Post) Count() int64 {
	var count int64
	database.DB.Model(models.Post{}).Count(&count)
	return count
}

func (p *Post) Create(m *models.Post) error {
	result := database.DB.Create(m)
	return result.Error
}
