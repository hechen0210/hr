package service

import (
	"hr/repository"
	"hr/util"
)

type IssueDemandService interface {
	GetList(page, perPage int, condition []repository.Condition) (info util.Page, err error)
	GetInfo(condition []repository.Condition) (info repository.IssueDemandInfo,err error)
}

type issueDemand struct {
	issueDemandRepository repository.IssueDemandRepository
}

func NewIssueDemandService() IssueDemandService {
	return &issueDemand{
		issueDemandRepository: repository.NewIssueDemandRepository(),
	}
}

func (id *issueDemand) GetList(page, perPage int, condition []repository.Condition) (info util.Page, err error) {
	total, err := id.issueDemandRepository.GetCount(condition)
	info = util.NewPage(total, perPage, page, nil)
	if err == nil {
		list,err := id.issueDemandRepository.GetList(page, perPage, condition)
		if err == nil {
			info.List = list	
		}
	}
	return info, err
}

func (id *issueDemand) GetInfo(conditon []repository.Condition) (info repository.IssueDemandInfo, err error) {
	return id.issueDemandRepository.GetInfo(conditon)
}