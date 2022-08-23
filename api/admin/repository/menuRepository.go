package repository

import (
	"hr/api/admin/model"
	"hr/config"

	"github.com/hechen0210/utils/helper"
	"github.com/jinzhu/gorm"
)

type MenuRepository interface {
	FindById(id int) (menu model.Menu, err error)
	GetList() (list []model.Menu, err error)
	Create(menu *model.Menu) error
	Update(id int, info *model.Menu) error
	UpdateShow(id, show int) error
	UpdateShowType(id, showType int) error
	Delete(id int) error
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository() MenuRepository {
	return &menuRepository{
		db: config.GetDbClient(),
	}
}

func (mr *menuRepository) FindById(id int) (menu model.Menu, err error) {
	err = mr.db.Where("id = ?", id).First(&menu).Error
	return
}

func (mr *menuRepository) GetList() (list []model.Menu, err error) {
	err = mr.db.Order("level").Find(&list).Error
	return
}

func (mr *menuRepository) Create(menu *model.Menu) (err error) {
	err = mr.db.Create(menu).Error
	return
}

func (mr *menuRepository) Update(id int, info *model.Menu) (err error) {
	data := helper.Struct2Map(info)
	delete(data, "id")
	delete(data, "created_at")
	err = mr.db.Model(&model.Menu{}).Where("id = ?", id).Update(data).Error
	return err
}

func (mr *menuRepository) UpdateShow(id, show int) (err error) {
	err = mr.db.Model(&model.Menu{}).Where("id = ?", id).Update("show", show).Error
	return
}

func (mr *menuRepository) UpdateShowType(id, showType int) (err error) {
	err = mr.db.Model(&model.Menu{}).Where("id = ?", id).Update("show_type", showType).Error
	return
}

func (mr *menuRepository) Delete(id int) (err error) {
	err = mr.db.Delete(&model.Menu{}, "id = ?", id).Error
	return
}
