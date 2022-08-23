package repository

import (
	"hr/config"
	"hr/model"
	"hr/util"

	"github.com/jinzhu/gorm"
)

type IssueDemandRepository interface {
	GetCount(condition []Condition) (count int, err error)
	GetList(page, perPage int, condition []Condition) (list []IssueDemandInfo, err error)
	GetInfo(condition []Condition) (info IssueDemandInfo, err error)
}

type issueDemand struct {
	db *gorm.DB
}

type IssueDemandInfo struct {
	model.DemandInfo
	model.IssueDemand
}

func NewIssueDemandRepository() IssueDemandRepository {
	return &issueDemand{
		db: config.GetDbClient(),
	}
}

func (id *issueDemand) GetCount(condition []Condition) (count int, err error) {
	err = id.db.Model(model.IssueDemand{}).Scopes(GetBy(condition)).Count(&count).Error
	return count, err
}

func (id *issueDemand) GetList(page, perPage int, condition []Condition) (list []IssueDemandInfo, err error) {
	err = id.query(condition).Scopes(util.Paginate(page, perPage)).Find(&list).Error
	return list, err
}

func (id *issueDemand) GetInfo(condition []Condition) (info IssueDemandInfo, err error) {
	err = id.query(condition).First(&info).Error
	return info, err
}

func (id *issueDemand) query(condition []Condition) *gorm.DB {
	return id.db.Model(model.IssueDemand{}).Scopes(GetBy(condition))
}
