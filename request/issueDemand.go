package request

import (
	"hr/repository"

	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type IssueDemandRequest struct {

}

type IssueDemandForm struct {
	Cdid int `form:"cdid" validate:"required|min:1"`
	Name string `form:"name" validate:"required"`
	Desc string `form:"desc" validate:"required"`
	Unit string `form:"unit" validate:"required|"`
}

func NewIssueDemandRequest() *IssueDemandRequest {
	return &IssueDemandRequest{}
}

func (idr *IssueDemandRequest) Search(map[string]string) []repository.Condition {
	return []repository.Condition{}
}

func (idr *IssueDemandRequest) GetFormData(ctx iris.Context) (data IssueDemandForm, err error) {
	err = ctx.ReadForm(&data)
	if err != nil {
		return data, err
	}
	validate := validate.Struct(data)
	return data, validate.ValidateE().OneError()
}