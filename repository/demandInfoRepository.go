package repository

import (
	"errors"
	"hr/config"
	"hr/model"
	"time"

	"github.com/jinzhu/gorm"
)

type DemandInfoRepository interface {
	GetInfo(condition []Condition) (*model.DemandInfo, error)
	Create(info model.DemandInfo, transaction *gorm.DB) error
	Update(info model.DemandInfo, transaction *gorm.DB) error
}

type demandInfoRepository struct {
	db *gorm.DB
}

func NewDemandInfoRepository() DemandInfoRepository {
	return &demandInfoRepository{
		db: config.GetDbClient(),
	}
}

// 获取详情
func (d *demandInfoRepository) GetInfo(condition []Condition) (info *model.DemandInfo, err error) {
	err = d.db.Scopes(GetBy(condition)).Find(&info).Error
	return info, err
}

// 创建需求详情
func (d *demandInfoRepository) Create(info model.DemandInfo, transaction *gorm.DB) (err error) {
	if info.Cdid == 0 {
		return errors.New("cdid is zero")
	}
	db := d.db
	if transaction != nil {
		db = transaction
	}
	info.CreatedAt = time.Now().Unix()
	err = db.Create(info).Error
	return err
}

// 更新需求详情
func (d *demandInfoRepository) Update(info model.DemandInfo, transaction *gorm.DB) (err error) {
	if info.Cdid == 0 {
		return errors.New("cdid is zero")
	}
	db := d.db
	if transaction != nil {
		db = transaction
	}
	err = db.Model(&model.DemandInfo{}).Update(map[string]interface{}{
		"type":        info.Type,
		"quantity":    info.Quantity,
		"sign_time":   info.SignTime,
		"education":   info.Education,
		"gender":      info.Gender,
		"age":         info.Age,
		"expire":      info.Expire,
		"expire_time": info.ExpireTime,
		"updated_at":  time.Now().Unix(),
	}).Where("cdid=?", info.Cdid).Error
	return err
}
