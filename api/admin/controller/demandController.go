package controller

import (
	"hr/request"
	"hr/response"
	"hr/service"
	"strconv"

	"github.com/kataras/iris/v12"
)

type DemandController struct {
	service service.DemandService
	request *request.DemandRequest
	response *response.Response
}

func NewDemandController() *DemandController {
	return &DemandController{
		service: service.NewDemandService(),
		request: request.NewDemandRequest(),
		response: response.NewResponse(),
	}
}

func (d *DemandController) GetList(ctx iris.Context) {
	page, perPage, condition := request.GetSearchParams(ctx, d.request.Search)
	list, err := d.service.GetList(page, perPage, condition)
	if err != nil {
		d.response.Fail(ctx)
		return
	}
	d.response.SetData(list).Success(ctx)
}

func (d *DemandController) GetInfo(ctx iris.Context) {
	did := ctx.URLParam("id")
	id, _ := strconv.Atoi(did)
	_, err := d.service.GetInfo(id)
	if err != nil {

		return
	}
}
