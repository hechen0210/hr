package controller

import (
	"hr/api/admin/request"
	"hr/api/admin/service"
	"hr/response"

	"github.com/kataras/iris/v12"
)

type MenuController struct {
	service  service.MenuService
	request  *request.MenuRequest
	response *response.Response
}

func NewMenuController() *MenuController {
	return &MenuController{
		service:  service.NewMenuService(),
		request:  request.NewMenuRequest(),
		response: response.NewResponse(),
	}
}

func (mc *MenuController) GetList(ctx iris.Context) {
	list, err := mc.service.GetList()
	if err != nil {
		mc.response.Fail(ctx)
		return
	}
	mc.response.SetData(list).Success(ctx)
}

func (mc *MenuController) Update(ctx iris.Context) {
	data, err := mc.request.GetMenuForm(ctx)
	if err != nil {

	}
	err = mc.service.CreateOrUpdate(data)
	if err != nil {
		mc.response.Fail(ctx)
		return
	}
	mc.response.Success(ctx)
}

func (mc *MenuController) Delete(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id == 0 {
		mc.response.Fail(ctx)
	}
	err := mc.service.Delete(id)
	if err != nil {
		mc.response.Fail(ctx)
	}
	mc.response.Success(ctx)
}
