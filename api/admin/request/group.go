package request

import (
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type GroupRequest struct {
}

type GroupForm struct {
	Id        int    `form:"id"`
	Name      string `form:"name" validate:"required"`
	Parent    int    `form:"parent"`
	Privilege string `form:"privilege"`
	Status    int    `form:"status"`
}

func NewGroupRequest() *GroupRequest {
	return &GroupRequest{}
}

func (gr *GroupRequest) GetFormData(ctx iris.Context) (data GroupForm, err error) {
	err = ctx.ReadForm(&data)
	if err != nil {
		return
	}
	validate := validate.Struct(data)
	return data, validate.ValidateE().OneError()
}
