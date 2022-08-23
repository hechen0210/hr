package request

import (
	"hr/repository"

	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type CompanyRequest struct {
}

type CompanyForm struct {
	Id       int    `form:"id"`
	Name     string `form:"name" validate:"required"`
	Account  string `form:"account" validate:"required|minLen:6|maxLen:20"`
	Password string `form:"password" validate:"required|minLen:6|maxLen:20"`
}

func NewCompanyRequest() *CompanyRequest {
	return &CompanyRequest{}
}

func (cr *CompanyRequest) Search(map[string]string) []repository.Condition {
	return []repository.Condition{}
}

func (cr *CompanyRequest) GetFormData(ctx iris.Context) (data CompanyForm, err error) {
	err = ctx.ReadForm(&data)
	if err != nil {
		return data, err	
	}
	validate := validate.Struct(data)
	return data, validate.ValidateE().OneError()
}
