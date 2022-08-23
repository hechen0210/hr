package service

import (
	"encoding/json"
	"hr/api/admin/model"
	"hr/api/admin/repository"
	"time"

	"github.com/jinzhu/gorm"
)

type PrivilegeService interface {
}

type privilegeService struct {
	repo repository.PrivilegeRepository
}

func NewPrivilegeService() PrivilegeService {
	return &privilegeService{
		repo: repository.NewPrivilegeRepository(),
	}
}

func (ps *privilegeService) GetByAdmin(admin int) {

}

func (ps *privilegeService) CreateOrUpdate(admin int, privilege []int) (err error) {
	privilegeStr, err := json.Marshal(privilege)
	if err != nil {
		return err
	}

	old, err := ps.repo.GetByAdmin(admin)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	newPrivilege := &model.Privilege{
		Privilege: string(privilegeStr),
		UpdatedAt: time.Now().Unix(),
	}
	if old.Id > 0 {
		return ps.repo.Update(old.Id, newPrivilege)
	}
	newPrivilege.Admin = admin
	newPrivilege.CreatedAt = time.Now().Unix()
	return ps.repo.Create(newPrivilege)
}
