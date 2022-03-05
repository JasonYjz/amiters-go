package api

import (
	"amiters-go/controller/resp"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

func (c *UserController) PostNew() *resp.JsonResult {
	return resp.SucData()
}

func (c *UserController) PostQuery() *resp.JsonResult {
	return resp.SucData()
}