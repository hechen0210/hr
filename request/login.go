package request

import (
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type LoginRequest struct {

}

type LoginForm struct {
	Account  string `form:"account" validate:"required|minLen:5|maxLen:20"`
	Password string `form:"password" validate:"required|minLen:5|maxLen:20"`
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func (lr *LoginRequest) GetFormData(ctx iris.Context) LoginForm {
	var login LoginForm
	ctx.ReadForm(&login)
	return login
}

func (lf *LoginForm) Validate() error {
	if err := validate.Struct(lf).ValidateE().OneError(); err != nil {
		return err
	}
	return nil
}