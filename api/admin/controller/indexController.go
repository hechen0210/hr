package controller

import (
	"hr/response"

	"github.com/kataras/iris/v12"
)

type IndexController struct {
	response *response.Response
}

func NewIndexController() *IndexController {
	return &IndexController{
		response: response.NewResponse(),
	}
}

func (ic *IndexController) List(ctx *iris.Context) {
	
}
