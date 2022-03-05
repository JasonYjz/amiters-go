package dao

import (
	"amiters-go/config"
	"amiters-go/model"
	"gorm.io/gorm"
)

var UserDao = &userDao{db: config.GetDB()}

type userDao struct {
	db *gorm.DB
}

func (u *userDao) Query(name string) *model.User {
	ret := &model.User{}
	if err := u.db.Take(ret, "name = ?", name).Error; err != nil {
		return nil
	}

	return ret
	//u.db.Where("name = ?", name).Get(&ret)
	//err = db.Model(&model.User{}).Where("id = ?", id).Updates(columns).Error
	//u.db.Model(&ret).Where("name = ?", name).
}

func (u *userDao) NewUser(usr *model.User) (err error) {
	return u.db.Create(usr).Error
}
