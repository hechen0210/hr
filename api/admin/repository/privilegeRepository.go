package repository

import (
	"hr/api/admin/model"
	"hr/config"

	"github.com/hechen0210/utils/helper"
	"github.com/jinzhu/gorm"
)

type PrivilegeRepository interface {
	GetAll() (list []model.Privilege,err error)
	GetById(id int) (info model.Privilege,err error)
	GetByAdmin(adminId int) (info model.Privilege,err error)
	Create(info *model.Privilege) (err error)
	Update(id int,info *model.Privilege) (err error)
}

type privilegeRepository struct {
	db *gorm.DB
}

func NewPrivilegeRepository() PrivilegeRepository {
	return &privilegeRepository{
		db: config.GetDbClient(),
	}
}

func (pr *privilegeRepository) GetAll() (list []model.Privilege, err error) {
	err = pr.db.Find(&list).Error
	return
}

func (pr *privilegeRepository) GetById(id int) (info model.Privilege, err error) {
	err = pr.db.Where("id = ?", id).First(info).Error
	return
}

func (pr *privilegeRepository) GetByAdmin(adminId int) (info model.Privilege, err error) {
	err = pr.db.Where("admin = ?", adminId).First(&info).Error
	return
}

func (pr *privilegeRepository) Create(info *model.Privilege) (err error) {
	err = pr.db.Create(&info).Error
	return
}

func (pr *privilegeRepository) Update(id int,info *model.Privilege) (err error) {
	data := helper.Struct2Map(info)
	delete(data, "id")
	delete(data, "created_at")
	err = pr.db.Model(&model.Privilege{}).Where("id = ?", id).Update(data).Error
	return
}