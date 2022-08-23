package repository

import (
	"hr/api/admin/model"
	"hr/config"

	"github.com/hechen0210/utils/helper"
	"github.com/jinzhu/gorm"
)

type RouteRepository interface {
	GetAll(routeType int) (list []model.Route, err error)
	Create(info *model.Route) (err error)
	Update(id int, info *model.Route) (err error)
}

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository() RouteRepository {
	return &routeRepository{
		db: config.GetDbClient(),
	}
}

func (rr *routeRepository) GetAll(routeType int) (list []model.Route, err error) {
	if routeType != 0 {
		err = rr.db.Where("type = ?", routeType).Find(&list).Error
		return
	}
	err = rr.db.Find(&list).Error
	return
}

func (rr *routeRepository) Create(info *model.Route) (err error) {
	err = rr.db.Create(info).Error
	return
}

func (rr *routeRepository) Update(id int, info *model.Route) (err error) {
	data := helper.Struct2Map(info)
	delete(data, "id")
	delete(data, "created_at")
	err = rr.db.Model(&model.Route{}).Where("id = ?", id).Update(data).Error
	return
}
