package controller

import (
	"hr/repository"
	"hr/request"
	"hr/response"
	"hr/service"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	loginService service.LoginService
	request      *request.LoginRequest
	response     *response.Response
}

func NewLoginController() *LoginController {
	return &LoginController{
		loginService: service.NewLoginService(repository.NewCompanyRepository()),
		request:      request.NewLoginRequest(),
		response:     response.NewResponse(),
	}
}

func (l *LoginController) Login(ctx iris.Context)  {
	data := l.request.LoginForm(ctx)
	if data.Validate() != nil {
		l.response.SetMessage(data.Validate().Error()).Fail(ctx)
		return
	}
	info, err := l.loginService.Login(data.Account, data.Password)
	if err != nil {
		l.response.SetMessage("登录失败").Fail(ctx)
		return
	}
	l.response.SetData(info).Success(ctx)
}

func (l *LoginController) Logout(ctx iris.Context) {

}
