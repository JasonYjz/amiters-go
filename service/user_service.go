package service

import (
	"amiters-go/dao"
	"amiters-go/model"
)

var UserService = new(userService)

type userService struct {
}

func (u *userService) Get(name string) *model.User {
	return dao.UserDao.Query(name)
}

func (u *userService) Create(name string) (err error) {
	user := &model.User{
		Name: name,
	}
	return dao.UserDao.NewUser(user)
}

