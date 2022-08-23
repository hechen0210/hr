package request

import (
	"hr/repository"

	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type AdminRequest struct {
}

type AdminForm struct {
	Id       int    `form:"id"`
	Account  string `form:"account" validate:"required|minLen:6|maxLen:20"`
	Name     string `form:"name" validate:"required"`
	Password string `form:"password" validate:"required|minLen:6|maxLen:20"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Group    int    `form:"group" validate:"required"`
	Status   int    `form:"status"`
}

func NewAdminRequest() *AdminRequest {
	return &AdminRequest{}
}

func (ar *AdminRequest) Search(map[string]string) []repository.Condition {
	return []repository.Condition{}
}

func (ar *AdminRequest) GetFormData(ctx iris.Context) (data AdminForm, err error) {
	err = ctx.ReadForm(&data)
	if err != nil {
		return data, err
	}
	validate := validate.Struct(data)
	return data, validate.ValidateE().OneError()
}
