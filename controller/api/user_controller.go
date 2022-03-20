package api

import (
	"amiters-go/controller/resp"
	"amiters-go/model"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

func (c *UserController) PostNew() *resp.JsonResult {
	fmt.Println("receive a request to create new user.")
	user := &model.User{}

	if body, err := c.Ctx.GetBody();err != nil {
		return resp.FailData()
	} else {
		err = json.Unmarshal(body, user)
		if err != nil {
			return resp.FailData()
		}

		fmt.Println(user)
	}


	//if err := service.UserService.Create(name); err != nil {
	//	return resp.FailData()
	//}
	return resp.SucData()
}

func (c *UserController) PostQuery() *resp.JsonResult {
	return resp.SucData()
}