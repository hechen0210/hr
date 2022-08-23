package repository

import (
	"hr/config"
	"hr/model"
	"hr/util"

	"github.com/jinzhu/gorm"
)

type CompanyRepository interface {
	FindByAccount(account string) *model.Company
	GetCount(condition []Condition) (count int, err error)
	GetList(page, perPage int, condition []Condition) (list []model.Company, err error)
	Create(info model.Company) (int, error)
	Update(id int, info model.Company) error
	Delete(id int) error
}

type companyRepository struct {
	db    *gorm.DB
}

func NewCompanyRepository() CompanyRepository {
	return &companyRepository{
		db:    config.GetDbClient(),
	}
}

func (cr *companyRepository) FindByAccount(account string) *model.Company {
	var company model.Company
	err := cr.db.Where("account=?", account).Find(&company)
	if err != nil {
		return nil
	}
	return &company
}

func (cr *companyRepository) GetCount(condition []Condition) (count int, err error) {
	err = cr.db.Model(model.Company{}).Scopes(GetBy(condition)).Count(&count).Error
	return count, err
}

func (cr *companyRepository) GetList(page, perPage int, condition []Condition) (list []model.Company, err error) {
	err = cr.db.Model(model.Company{}).Scopes(GetBy(condition)).Scopes(util.Paginate(page, perPage)).Find(&list).Error
	return list, err
}

func (cr *companyRepository) Create(info model.Company) (int, error) {
	db := cr.db.Begin()
	err := db.Create(&info).Error
	if err == nil {
		companyInfo := &model.CompanyInfo{
			Cid: info.Id,
		}
		err = db.Create(&companyInfo).Error
	}
	if err != nil {
		db.Rollback()
		return 0, err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, err
	}
	return info.Id, nil
}

func (cr *companyRepository) Update(id int,info model.Company) error {
	return cr.db.Model(&model.Company{}).Where("id=?", id).Update(info).Error
}

func (cr *companyRepository) Delete(id int) error {
	return cr.db.Model(&model.Company{}).Where("id=?", id).Update("flag", -1).Error
}
