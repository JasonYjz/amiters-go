package api

import (
	"amiters-go/controller/resp"
	"amiters-go/service"
	"fmt"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

func (c *UserController) PostNew() *resp.JsonResult {
	fmt.Println("receive a request to create new user.")
	var (
		name = c.Ctx.PostValueTrim("name")
	)

	if err := service.UserService.Create(name); err != nil {
		return resp.FailData()
	}
	return resp.SucData()
}

func (c *UserController) PostQuery() *resp.JsonResult {
	return resp.SucData()
}