package service

import (
	"hr/api/admin/model"
	"hr/api/admin/repository"
	"hr/api/admin/request"
	"time"
)

type GroupService interface {
	GetList() (list []model.Group, err error)
	CreateOrUpdate(info request.GroupForm) (err error)
	ChangeStatus(id int, status int) (err error)
	Delete(id int) (err error)
}

type groupService struct {
	repo repository.GroupRepository
}

func NewGroupService() GroupService {
	return &groupService{
		repo: repository.NewGroupRepository(),
	}
}

func (gs *groupService) GetList() (list []model.Group, err error) {
	return gs.repo.GetAll()
}

func (gs *groupService) CreateOrUpdate(info request.GroupForm) (err error) {
	group := &model.Group{
		Name:      info.Name,
		Parent:    info.Parent,
		Privilege: info.Privilege,
		UpdatedAt: time.Now().Unix(),
	}
	if info.Id > 0 {
		return gs.repo.Update(info.Id, group)
	}
	group.CreatedAt = time.Now().Unix()
	return gs.repo.Create(group)
}

func (gs *groupService) ChangeStatus(id int, status int) (err error) {
	return gs.repo.ChangeStatus(id, status)
}

func (gs *groupService) Delete(id int) (err error) {
	return gs.repo.Delete(id)
}
