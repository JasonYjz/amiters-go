package service

import (
	"amiters-go/dao"
	"amiters-go/model"
	"fmt"
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
	fmt.Println(user)
	return dao.UserDao.NewUser(user)
}

