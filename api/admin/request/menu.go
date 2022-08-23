package request

import (
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type MenuRequest struct {
}

type MenuForm struct {
	Id       int    `form:"id"`
	Name     string `form:"name" validate:"required"`
	Parent   int    `form:"parent"`
	Show     int    `form:"show"`
	ShowType int    `form:"show_type"`
	Url      string `form:"url" validate:"required"`
	Icon     string `form:"icon"`
	Sort     int    `form:"sort"`
	Api      string `form:"api"`
}

func NewMenuRequest() *MenuRequest {
	return &MenuRequest{}
}

func (mr *MenuRequest) GetMenuForm(ctx iris.Context) (data MenuForm, err error) {
	err = ctx.ReadForm(&data)
	if err != nil {
		return
	}
	validate := validate.Struct(data)
	return data, validate.ValidateE().OneError()
}
