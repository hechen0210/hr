package repository

import (
	"hr/config"
	"hr/model"
	"hr/util"

	"github.com/jinzhu/gorm"
)

type CompanySalaryInfo struct {
}

type CompanySalaryRepository interface {
	GetCount(condition []Condition) (int, error)
	GetList(page, perPage int, condition []Condition) ([]CompanySalaryInfo, error)
	GetInfo(id int) (CompanySalaryInfo, error)
}

type companySalaryRepository struct {
	db *gorm.DB
}

func NewCompanySalaryRepo() CompanySalaryRepository {
	return &companySalaryRepository{
		db: config.GetDbClient(),
	}
}

func (cs *companySalaryRepository) GetCount(condition []Condition) (count int, err error) {
	err = cs.db.Model(model.CompanySalary{}).Scopes(GetBy(condition)).Count(&count).Error
	return count, err
}

func (cs *companySalaryRepository) GetList(page, perPage int, condition []Condition) (data []CompanySalaryInfo, err error) {
	err = cs.db.Table(model.CompanySalary{}.TableName()).Scopes(GetBy(condition)).Scopes(util.Paginate(page, perPage)).Find(&data).Error
	return data, err
}

func (cs *companySalaryRepository) GetInfo(id int) (info CompanySalaryInfo, err error) {
	return info, err
}
