package service

import (
	"errors"
	"hr/config"
	"hr/model"
	"hr/repository"
	"hr/request"
	"hr/util"
	"time"

	"github.com/jinzhu/gorm"
)

type AdminService interface {
	GetByName(name string) (info model.Admin, err error)
	GetById(id int) (info model.Admin, err error)
	GetList(page, perPage int, condition []repository.Condition) (info util.Page, err error)
	CreateOrUpdate(info request.AdminForm) (err error)
	CheckUserExist(info request.AdminForm) error
}

type adminService struct {
	repo repository.AdminRepository
}

func NewAdminService() AdminService {
	return &adminService{
		repo: repository.NewAdminRepository(),
	}
}

func (as *adminService) GetByName(name string) (info model.Admin, err error) {
	return as.repo.FindByName(name)
}

func (as *adminService) GetById(id int) (info model.Admin, err error) {
	return as.repo.FindById(id)
}

func (as *adminService) GetList(page, perPage int, condition []repository.Condition) (info util.Page, err error) {
	count, err := as.repo.GetCount(condition)
	info = util.NewPage(count, perPage, page, nil)
	if err == nil {
		list, err := as.repo.GetList(page, perPage, condition)
		if err == nil {
			info.List = list
		}
	}
	return info, err
}

func (as *adminService) CreateOrUpdate(info request.AdminForm) (err error) {
	if err = as.CheckUserExist(info); err != nil {
		return
	}
	user := model.Admin{
		Account:   info.Account,
		Name:      info.Name,
		Phone:     info.Phone,
		Email:     info.Email,
		Group:     info.Group,
		Status:    info.Status,
		UpdatedAt: time.Now().Unix(),
	}
	if info.Id == 0 || (info.Id > 0 && info.Password != "") {
		password, err := util.BcryptPassword(info.Password)
		if err != nil {
			return err
		}
		user.Password = password
	}
	if info.Id > 0 {
		return as.repo.Update(info.Id, user)
	}
	user.CreatedAt = time.Now().Unix()
	return as.repo.Create(user)
}

func (as *adminService) CheckUserExist(info request.AdminForm) (err error) {
	if err = as.checkAccountExist(info.Id, info.Account); err != nil {
		return err
	}
	if err = as.checkNameExist(info.Id, info.Name); err != nil {
		return err
	}
	if err = as.checkPhoneExist(info.Id, info.Phone); err != nil {
		return err
	}
	if err = as.checkEmailExist(info.Id, info.Email); err != nil {
		return err
	}
	return nil
}

func (as *adminService) checkAccountExist(id int, account string) error {
	user, err := as.repo.FindByAcount(account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	if user.Id == id {
		return nil
	}
	return errors.New(config.ACCOUNT_EXIST)
}

func (as *adminService) checkNameExist(id int, name string) error {
	user, err := as.repo.FindByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	if user.Id == id {
		return nil
	}
	return errors.New(config.NAME_EXIST)
}

func (as *adminService) checkPhoneExist(id int, phone string) error {
	if phone == "" {
		return nil
	}
	user, err := as.repo.FindByPhone(phone)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	if user.Id == id {
		return nil
	}
	return errors.New(config.PHONE_EXIST)
}

func (as *adminService) checkEmailExist(id int, email string) error {
	if email == "" {
		return nil
	}
	user, err := as.repo.FindByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	if user.Id == id {
		return nil
	}
	return errors.New(config.EMAIL_EXIST)
}
