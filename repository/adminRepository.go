package repository

import (
	"hr/config"
	"hr/model"
	"hr/util"

	"github.com/jinzhu/gorm"
)

type AdminRepository interface {
	FindByName(name string) (info model.Admin, err error)
	FindById(id int) (info model.Admin, err error)
	FindByAcount(account string) (info model.Admin, err error)
	FindByPhone(phone string) (info model.Admin, err error)
	FindByEmail(email string) (info model.Admin, err error)
	GetCount(condition []Condition) (count int, err error)
	GetList(page, pageSize int, condition []Condition) (list []model.Admin, err error)
	Create(info model.Admin) (err error)
	Update(id int, info model.Admin) (err error)
	Delete(id int) (err error)
	Frzee(id int) (err error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository() AdminRepository {
	return &adminRepository{
		db: config.GetDbClient(),
	}
}

func (ar *adminRepository) FindByName(name string) (info model.Admin, err error) {
	err = ar.db.Where("name = ?", name).First(&info).Error
	return
}

func (ar *adminRepository) FindById(id int) (info model.Admin, err error) {
	err = ar.db.Where("id = ?", id).First(&info).Error
	return
}

func (ar *adminRepository) FindByAcount(account string) (info model.Admin, err error) {
	err = ar.db.Debug().Where("account = ?", account).First(&info).Error
	return
}

func (ar *adminRepository) FindByPhone(phone string) (info model.Admin, err error) {
	err = ar.db.Where("phone = ?", phone).First(&info).Error
	return
}

func (ar *adminRepository) FindByEmail(email string) (info model.Admin, err error) {
	err = ar.db.Where("email = ?", email).First(&info).Error
	return
}

func (ar *adminRepository) GetCount(condition []Condition) (count int, err error) {
	err = ar.db.Model(&model.Admin{}).Scopes(GetBy(condition)).Count(&count).Error
	return
}

func (ar *adminRepository) GetList(page, pageSize int, condition []Condition) (list []model.Admin, err error) {
	err = ar.db.Model(&model.Admin{}).Scopes(GetBy(condition)).Scopes(util.Paginate(page, pageSize)).Find(&list).Error
	return
}

func (ar *adminRepository) Create(info model.Admin) (err error) {
	err = ar.db.Create(&info).Error
	return
}

func (ar *adminRepository) Update(id int, info model.Admin) (err error) {
	data := make(map[string]interface{})
	if info.Password != "" {
		data["password"] = info.Password
	}
	data["account"] = info.Account
	data["name"] = info.Name
	data["phone"] = info.Phone
	data["email"] = info.Email
	data["group"] = info.Group
	data["status"] = info.Status
	data["updated_at"] = info.UpdatedAt
	err = ar.db.Debug().Model(&model.Admin{}).Where("id = ?", id).Updates(data).Error
	return
}

func (ar *adminRepository) Delete(id int) (err error) {
	err = ar.db.Where("id = ?", id).Delete(&model.Admin{}).Error
	return
}

func (ar *adminRepository) Frzee(id int) (err error) {
	err = ar.db.Model(&model.Admin{}).Where("id = ?", id).Update("status", 0).Error
	return
}
