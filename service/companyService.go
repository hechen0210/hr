package service

import (
	"hr/model"
	"hr/repository"
	"hr/request"
	"hr/util"
	"time"
)

type CompanyService interface {
	GetList(page, perPage int, condition []repository.Condition) (util.Page, error)
	GetInfo(id int)
	Create() error
	EditAccount(data request.CompanyForm) error
	EditInfo() error
	Delete(id int) error
}

type companyService struct {
	companyRepo repository.CompanyRepository
}

func NewCompanyService() CompanyService {
	return &companyService{
		companyRepo: repository.NewCompanyRepository(),
	}
}

func (cs *companyService) GetList(page, perPage int, condition []repository.Condition) (util.Page, error) {
	count, err := cs.companyRepo.GetCount(condition)
	info := util.NewPage(count, perPage, page, nil)
	if err == nil {
		list, err := cs.companyRepo.GetList(page, perPage, condition)
		if err == nil {
			info.List = list
		}
	}
	return info, err
}

func (cs *companyService) GetInfo(id int) {

}

func (cs *companyService) Create() error {
	return nil
}

func (cs *companyService) EditAccount(data request.CompanyForm) error {
	company := model.Company{
		Account:   data.Account,
		Name:      data.Name,
		UpdatedAt: time.Now().Unix(),
	}
	if data.Password != "" {
		password,err := util.BcryptPassword(data.Password)
		if err != nil {
			return err
		}
		company.Password = password
	}
	return cs.companyRepo.Update(data.Id, company)
}

func (cs *companyService) EditInfo() error {
	return nil
}

func (cs *companyService) Delete(id int) error {
	return cs.companyRepo.Delete(id)
}
