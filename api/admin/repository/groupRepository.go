package repository

import (
	"hr/api/admin/model"
	"hr/config"

	"github.com/hechen0210/utils/helper"
	"github.com/jinzhu/gorm"
)

type GroupRepository interface {
	GetAll() (list []model.Group, err error)
	GetById(id int) (info model.Group, err error)
	Create(info *model.Group) (err error)
	Update(id int, info *model.Group) (err error)
	ChangeStatus(id, status int) (err error)
	Delete (id int) (err error)
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository() GroupRepository {
	return &groupRepository{
		db: config.GetDbClient(),
	}
}

func (gr *groupRepository) GetAll() (list []model.Group, err error) {
	err = gr.db.Find(&list).Error
	return
}

func (gr *groupRepository) GetById(id int) (info model.Group, err error) {
	err = gr.db.Where("id = ?", id).First(info).Error
	return
}

func (gr *groupRepository) Create(info *model.Group) (err error) {
	err = gr.db.Create(info).Error
	return
}

func (gr *groupRepository) Update(id int, info *model.Group) (err error) {
	data := helper.Struct2Map(info)
	delete(data, "id")
	delete(data, "created_at")
	err = gr.db.Model(&model.Group{}).Where("id = ?", id).Update(data).Error
	return
}

func (gr *groupRepository) ChangeStatus(id, status int) (err error) {
	err = gr.db.Model(&model.Group{}).Where("id = ?", id).Update("status", status).Error
	return
}

func (gr *groupRepository) Delete(id int) (err error) {
	err = gr.db.Where("id = ?", id).Delete(&model.Group{}).Error
	return
}