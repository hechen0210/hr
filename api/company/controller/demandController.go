package controller

import (
	"hr/config"
	"hr/request"
	"hr/response"
	"hr/service"
	"strconv"

	"github.com/kataras/iris/v12"
)

type DemandController struct {
	DemandService service.DemandService
	request       *request.DemandRequest
	response      *response.Response
}

func NewDemandController() *DemandController {
	return &DemandController{
		DemandService: service.NewDemandService(),
		request:       request.NewDemandRequest(),
		response:      response.NewResponse(),
	}
}

// Get Serves
// Method GET
// Resource: http://127.0.0.1:8080/demand
func (d *DemandController) GetList(ctx iris.Context) {
	page, perPage, condition := request.GetSearchParams(ctx, d.request.Search)
	list, err := d.DemandService.GetList(page, perPage, condition)
	if err != nil {
		d.response.Fail(ctx)
		return
	}
	d.response.SetData(list).Success(ctx)
}

func (d *DemandController) GetInfo(ctx iris.Context) {
	did := ctx.URLParam("id")
	id, _ := strconv.Atoi(did)
	_, err := d.DemandService.GetInfo(id)
	if err != nil {

		return
	}

}

// PostEdit Serves
// Method POST
// Resource: http://127.0.0.1:8080/demand/edit
func (d *DemandController) Edit(ctx iris.Context) {
	data, err := d.request.FormData(ctx)
	if err != nil {
		d.response.SetCode(config.Fail).Fail(ctx)
		return
	}
	_, err = d.DemandService.Create(data)
	if err != nil {
		return
	}
}

// Delete Serves
// Method Delete
// Resource: http://127.0.0.1:8080/delete
func (d *DemandController) Delete(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id == 0 {
		return
	}
	d.DemandService.Delete(id)
}
