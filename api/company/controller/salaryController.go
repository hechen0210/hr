package controller

import (
	"fmt"
	"hr/request"
	"hr/service"

	"github.com/kataras/iris/v12"
)

type SalaryController struct {
	SalaryService service.CompanySalaryService
	request       request.CompanySalaryRequest
}

func NewSalaryController() *SalaryController {
	return &SalaryController{
		SalaryService: service.NewCompanySalaryService(),
		request:       request.NewCompanySalaryRequest(),
	}
}

func (s *SalaryController) GetList(ctx iris.Context) {
	page, perPage, condition := request.GetSearchParams(ctx, s.request.Search)
	info, err := s.SalaryService.GetList(page, perPage, condition)
	if err != nil {
		return
	}
	fmt.Println(info)
}

func (s *SalaryController) GetInfo(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id == 0 {
		return
	}
	info, err := s.SalaryService.GetInfo(id)
	if err != nil {
		return
	}
	fmt.Println(info)
}
