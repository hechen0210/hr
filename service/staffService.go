package service

import (
	"hr/repository"
	"hr/util"
)

type StaffService interface {
	GetList(page, perPage int, condition []repository.Condition) (info util.Page, err error)
	GetInfo(id int) (info map[string]interface{}, err error)
}

type staffService struct {
	staffRepo repository.StaffRepository
}

func NewStaffService() StaffService {
	return &staffService{
		staffRepo: repository.NewStaffRepository(),
	}
}

func (s *staffService) GetList(page, perPage int, condition []repository.Condition) (info util.Page, err error) {
	total, err := s.staffRepo.GetCount(condition)
	if err != nil {
		return info, err
	}
	info = util.NewPage(total, perPage, page, []repository.StaffInfo{})
	if total > 0 {
		list, err := s.staffRepo.GetList(page, perPage, condition)
		if err != nil {
			return info, err
		}
		info.List = list
	}
	return info, err
}

func (s *staffService) GetInfo(id int) (info map[string]interface{}, err error) {
	data, err := s.staffRepo.GetInfo(id)
	if err != nil {
		return info, err
	}
	other := map[string]interface{}{}
	return s.formatData(data, other), err
}

func (s *staffService) formatData(info repository.StaffInfo, otherInfo map[string]interface{}) map[string]interface{} {
	parse := (&Formater{}).Format(info, otherInfo)
	return parse.Result
}
