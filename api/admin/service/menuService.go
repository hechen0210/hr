package service

import (
	"hr/api/admin/model"
	"hr/api/admin/repository"
	"hr/api/admin/request"
	"time"
)

type MenuService interface {
	GetList() (list []Menu, err error)
	CreateOrUpdate(data request.MenuForm) (err error)
	ChangeShow(id int) (err error)
	ChangeShowType(id, showType int) (err error)
	Delete(id int) (err error)
}

type menuService struct {
	repo repository.MenuRepository
}

type Menu struct {
	model.Menu
	Son []Menu
}

func NewMenuService() MenuService {
	return &menuService{
		repo: repository.NewMenuRepository(),
	}
}

func (ms *menuService) GetList() (list []Menu, err error) {
	menus, err := ms.repo.GetList()
	if err != nil {
		return list, err
	}
	son := make(map[int][]int)
	for _, menu := range menus {
		if menu.Level != 1 {
			son[menu.Parent] = append(son[menu.Parent], menu.Id)
		}
	}
	return ms.menuTree(menus, son), err
}

func (ms *menuService) menuTree(menus []model.Menu, sons map[int][]int) (list []Menu) {
	for index, menu := range menus {
		if _, ok := sons[menu.Id]; ok {
			list = ms.menuTree(menus[index+1:], sons)
		} else {
			list = append(list, Menu{Menu: menu})
		}
	}
	return list
}

func (ms *menuService) CreateOrUpdate(data request.MenuForm) (err error) {
	parent := model.Menu{}
	if data.Parent != 0 {
		parent, err = ms.repo.FindById(data.Parent)
		if err != nil {
			return
		}
	}
	menu := &model.Menu{
		Name:      data.Name,
		Parent:    data.Parent,
		Show:      data.Show,
		ShowType:  data.ShowType,
		Url:       data.Url,
		Icon:      data.Icon,
		Sort:      data.Sort,
		Level:     parent.Level + 1,
		Api:       data.Api,
		UpdatedAt: time.Now().Unix(),
	}
	if data.Id > 0 {
		return ms.repo.Update(data.Id, menu)
	}
	menu.CreatedAt = time.Now().Unix()
	return ms.repo.Create(menu)
}

func (ms *menuService) ChangeShow(id int) (err error) {
	menu, err := ms.repo.FindById(id)
	if err != nil {
		return
	}
	show := 0
	if menu.Show == 0 {
		show = 1
	}
	return ms.repo.UpdateShow(id, show)
}

func (ms *menuService) ChangeShowType(id, showType int) (err error) {
	return ms.repo.UpdateShowType(id, showType)
}

func (ms *menuService) Delete(id int) (err error) {
	return ms.repo.Delete(id)
}
