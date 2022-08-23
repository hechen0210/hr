package service

import (
	"hr/repository"
	"hr/util"
)

type CompanySalaryService interface {
	GetList(page, perPage int, condition []repository.Condition) (info util.Page, err error)
	GetInfo(id int) (info map[string]interface{}, err error)
}

type companySalaryService struct {
	companySalaryRepo repository.CompanySalaryRepository
}

func NewCompanySalaryService() CompanySalaryService {
	return &companySalaryService{
		companySalaryRepo: repository.NewCompanySalaryRepo(),
	}
}

func (css *companySalaryService) GetList(page, perPage int, condition []repository.Condition) (info util.Page, err error) {
	count, err := css.companySalaryRepo.GetCount(condition)
	if err != nil {
		return info, err
	}
	info = util.NewPage(count, perPage, page, []repository.CompanySalaryInfo{})
	if err != nil {
		list, err := css.companySalaryRepo.GetList(page, perPage, condition)
		if err != nil {
			return info, err
		}
		info.List = list
	}
	return info, err
}

func (css *companySalaryService) GetInfo(id int) (info map[string]interface{}, err error) {
	return info, err
}
