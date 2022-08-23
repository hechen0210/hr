package repository

import (
	"hr/config"
	"hr/model"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type DemandInfo struct {
	model.CompanyDemand
	model.DemandInfo
}

type DemandRepository interface {
	GetCount(condition []Condition) (int, error)
	GetList(page, pageSize int, condition []Condition) ([]DemandInfo, error)
	GetInfo(id int) (DemandInfo, error)
	Delete(id int) error
	Create(base model.CompanyDemand, info model.DemandInfo) (int, error)
	Update(id int, base model.CompanyDemand, info model.DemandInfo) error
	ChangeFlag(id, flag int) error
	Published(id int) error
}

type demandRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewDemandRepository() DemandRepository {
	return &demandRepository{
		db:    config.GetDbClient(),
		redis: config.GetRedis(),
	}
}

// 获取总数
func (d *demandRepository) GetCount(condition []Condition) (count int, err error) {
	err = d.db.Model(model.CompanyDemand{}).Scopes(GetBy(condition)).Count(&count).Error
	return count, err
}

// 获取列表
func (d *demandRepository) GetList(page, pageSize int, condition []Condition) (list []DemandInfo, err error) {
	err = d.db.Select("cd.*,di.*").
		Table(model.CompanyDemand{}.TableName() + " cd").
		Joins("left join " + model.DemandInfo{}.TableName() + " di on cd.id = di.cdid").
		Scopes(GetBy(condition)).
		Find(&list).Error
	return list, err
}

// 获取需求详情
func (d *demandRepository) GetInfo(id int) (info DemandInfo, err error) {
	err = d.db.Select("cd.*,di.*").
		Table(model.CompanyDemand{}.TableName() + " cd").
		Joins("left join " + model.DemandInfo{}.TableName() + " di on cd.id = di.cdid").
		First(&info).Error
	return info, err
}

// 删除需求
func (d *demandRepository) Delete(id int) (err error) {
	return d.db.Model(&model.CompanyDemand{}).Update("flag = -1").Error
}

// 创建需求
func (d *demandRepository) Create(base model.CompanyDemand, info model.DemandInfo) (id int, err error) {
	db := d.db.Begin()
	base.CreatedAt = time.Now().Unix()
	base.UpdatedAt = time.Now().Unix()
	err = db.Create(&base).Error
	if err == nil {
		info.Cdid = base.Id
		demandInfo := NewDemandInfoRepository()
		err = demandInfo.Create(info, db)
	}
	if err != nil {
		db.Rollback()
		return 0, err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, err
	}
	return base.Id, nil
}

// 更新需求
func (d *demandRepository) Update(id int, base model.CompanyDemand, info model.DemandInfo) (err error) {
	var baseInfo model.CompanyDemand
	err = d.db.Where("id = ?", id).Find(&baseInfo).Error
	if err != nil {
		return err
	}
	baseInfo.Name = base.Name
	baseInfo.Desc = base.Desc
	baseInfo.Unit = base.Unit
	base.UpdatedAt = time.Now().Unix()
	db := d.db.Begin()
	err = db.Save(&info).Error
	if err == nil {
		demandInfo := NewDemandInfoRepository()
		err = demandInfo.Update(info, db)
	}
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return err
}

// 修改需求状态
func (d *demandRepository) ChangeFlag(id, flag int) (err error) {
	err = d.db.Model(&model.CompanyDemand{}).Update("flag=?", flag).Where("id=?", id).Error
	return err
}

// 修改需求在平台的状态
func (d *demandRepository) Published(id int) (err error) {
	err = d.db.Model(&model.CompanyDemand{}).Update("published=1").Where("id=?", id).Error
	return err
}
