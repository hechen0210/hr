package request

import (
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type RouteRequest struct {
}

type RouteForm struct {
	Id   int    `form:"id"`
	Name string `form:"name" validate:"required"`
	Type int    `form:"type" validate:"required"`
	Url  string `form:"url" validate:"required"`
}

func NewRouteRequest() *RouteRequest {
	return &RouteRequest{}
}

func (rr *RouteRequest) GetRouteForm(ctx iris.Context) (data RouteForm, err error) {
	err = ctx.ReadForm(&data)
	if err != nil {
		return
	}
	validate := validate.Struct(data)
	return data, validate.ValidateE().OneError()
}
