package repository

import (
	"hr/config"
	"hr/model"

	"github.com/jinzhu/gorm"
)

type StaffInfo struct {
	model.CompanyStaff
	model.UserProfile
	model.CompanyDemand
}

type StaffRepository interface {
	GetCount(condition []Condition) (int, error)
	GetList(page, pageSize int, condition []Condition) ([]StaffInfo, error)
	GetInfo(id int) (StaffInfo, error)
}

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository() StaffRepository {
	return &staffRepository{
		db: config.GetDbClient(),
	}
}

func (s *staffRepository) GetCount(condition []Condition) (count int, err error) {
	err = s.db.Model(model.CompanyStaff{}).Scopes(GetBy(condition)).Count(&count).Error
	return count, err
}

func (s *staffRepository) GetList(page, perPage int, condition []Condition) (data []StaffInfo, err error) {
	err = s.query(condition).Find(&data).Error
	return data, err
}

func (s *staffRepository) GetInfo(id int) (data StaffInfo, err error) {
	err = s.query([]Condition{{"id", "=", id}}).Find(&data).Error
	return data, err
}

func (s *staffRepository) query(condition []Condition) *gorm.DB {
	return s.db.Select("cs.*,up.*,cd.*").Table(model.CompanyStaff{}.TableName() + " cs").
		Joins("left join " + model.UserProfile{}.TableName() + " up on cs.uid = up.uid").
		Joins("left join " + model.CompanyDemand{}.TableName() + " cd on cs.id = cd.csid").
		Scopes(GetBy(condition))
}
