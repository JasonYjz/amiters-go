package dao

import (
	"amiters-go/model"
	"gorm.io/gorm"
)

type postDao struct {
	DB *gorm.DB
}

func (p *postDao) NewPost(post *model.Post) (err error) {
	return p.DB.Create(post).Error
}

func (p *postDao) QueryById(id int) *model.Post {
	ret := &model.Post{}
	if err := p.DB.Take(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
