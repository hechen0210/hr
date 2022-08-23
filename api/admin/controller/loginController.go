package controller

import (
	"fmt"
	"hr/repository"
	"hr/request"
	"hr/response"
	"hr/service"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	service  service.LoginService
	request  *request.LoginRequest
	response *response.Response
}

func NewLoginController() *LoginController {
	return &LoginController{
		service:  service.NewAdminLoginService(repository.NewAdminRepository()),
		request:  request.NewLoginRequest(),
		response: response.NewResponse(),
	}
}

func (lc *LoginController) Login(ctx iris.Context) {
	loginForm := lc.request.GetFormData(ctx)
	if err := loginForm.Validate(); err != nil {
		fmt.Println(err)
		lc.response.Fail(ctx)
		return
	}
	token, err := lc.service.Login(loginForm.Account, loginForm.Password)
	if err != nil {
		fmt.Println(err)
		lc.response.Fail(ctx)
		return
	}
	lc.response.SetData(map[string]string{
		"token": token,
	}).Success(ctx)
}

func (lc *LoginController) Logout(ctx iris.Context) {

}
