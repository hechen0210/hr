package controller

import (
	"hr/request"
	"hr/response"
	"hr/service"

	"github.com/kataras/iris/v12"
)

type CompanyController struct {
	service  service.CompanyService
	request  *request.CompanyRequest
	response *response.Response
}

func NewCompanyController() *CompanyController {
	return &CompanyController{
		service:  service.NewCompanyService(),
		request:  request.NewCompanyRequest(),
		response: response.NewResponse(),
	}
}

func (cc *CompanyController) GetList(ctx iris.Context) {
	page, perPage, condition := request.GetSearchParams(ctx, cc.request.Search)
	list, err := cc.service.GetList(page, perPage, condition)
	if err != nil {
		cc.response.Fail(ctx)
	}
	cc.response.SetData(list).Success(ctx)
}

func (cc *CompanyController) GetInfo(ctx iris.Context) {

}

func (cc *CompanyController) Update(ctx iris.Context) {
	data,err := cc.request.FormData(ctx)
	if err != nil {
		cc.response.Fail(ctx)
	}
	if data.Id > 0 {
		err = cc.service.EditAccount(data)
		if err != nil {
			cc.response.Fail(ctx)
		}
	}
	err = cc.service.Create()
	if err != nil {
		cc.response.Fail(ctx)
	}
	cc.response.Success(ctx)
}

func (cc *CompanyController) Delete(ctx iris.Context) {
	id := ctx.Values().GetIntDefault("id",0)
	if id == 0 {
		cc.response.Fail(ctx)
	}
	err := cc.service.Delete(id)
	if err != nil {
		cc.response.Fail(ctx)
	}
	cc.response.Success(ctx)
}