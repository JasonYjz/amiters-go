package api

import (
	"amiters-go/controller/resp"
	"fmt"
	"github.com/kataras/iris/v12"
	"strconv"
)

type PostController struct {
	Ctx iris.Context
}

// PostCreate 发表文章
func (c *PostController) PostCreate() *resp.JsonResult {
	fmt.Println(c.Ctx)
	return resp.JsonData(200, "OK")
}

func (c *PostController) GetBy(postId int) *resp.JsonResult {
	return resp.JsonData(200, "postId:" + strconv.Itoa(postId))
}
