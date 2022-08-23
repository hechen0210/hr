package controller

import (
	"fmt"
	"hr/repository"
	"hr/request"
	"hr/response"
	"hr/service"

	"github.com/kataras/iris/v12"
)

type IssueDemandController struct {
	service  service.IssueDemandService
	request  *request.IssueDemandRequest
	response *response.Response
}

func NewIssueDemandController() *IssueDemandController {
	return &IssueDemandController{
		service:  service.NewIssueDemandService(),
		request:  request.NewIssueDemandRequest(),
		response: response.NewResponse(),
	}
}

func (ic *IssueDemandController) GetList(ctx iris.Context) {
	page, perPage, condition := request.GetSearchParams(ctx, ic.request.Search)
	list, err := ic.service.GetList(page, perPage, condition)
	if err != nil {
		ic.response.Fail(ctx)
		return
	}
	ic.response.SetData(list).Success(ctx)
}

func (ic *IssueDemandController) GetInfo(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id == 0 {
		ic.response.Fail(ctx)
	}
	info, err := ic.service.GetInfo([]repository.Condition{{
		Name:  "id",
		Value: id,
	}})
	if err != nil {
		ic.response.Fail(ctx)
		return
	}
	ic.response.SetData(info).Success(ctx)
}

func (ic *IssueDemandController) Update(ctx iris.Context) {
	data,err := ic.request.FormData(ctx)
	if err != nil {
		ic.response.Fail(ctx)
	}
	fmt.Println(data)
}

func (ic *IssueDemandController) Delete(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id == 0 {
		ic.response.Fail(ctx)
	}
}
