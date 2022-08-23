package controller

import (
	"fmt"
	"hr/request"
	"hr/service"

	"github.com/kataras/iris/v12"
)

type StaffController struct {
	staffService service.StaffService
	request      request.StaffRequest
}

func NewStaffController() *StaffController {
	return &StaffController{
		staffService: service.NewStaffService(),
		request:      request.NewStaffRequest(),
	}
}

func (s *StaffController) GetList(ctx iris.Context) {
	page, perPage, condition := request.GetSearchParams(ctx, s.request.Search)
	info, err := s.staffService.GetList(page, perPage, condition)
	if err != nil {
		return
	}
	fmt.Println(info)
}

func (s *StaffController) GetInfo(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id == 0 {
		return
	}
	info, err := s.staffService.GetInfo(id)
	if err != nil {
		return
	}
	fmt.Println(info)
}
