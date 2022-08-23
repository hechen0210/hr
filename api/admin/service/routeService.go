package service

import (
	"hr/api/admin/model"
	"hr/api/admin/repository"
	"hr/api/admin/request"
	"time"
)

type RouteService interface {
	GetList(routeType int) (list []model.Route, err error)
	CreateOrUpdate(info request.RouteForm) (err error)
}

type routeService struct {
	repo repository.RouteRepository
}

func NewRouteService() RouteService {
	return &routeService{
		repo: repository.NewRouteRepository(),
	}
}

func (rs *routeService) GetList(routeType int) (list []model.Route, err error) {
	list, err = rs.repo.GetAll(routeType)
	return
}

func (rs *routeService) CreateOrUpdate(info request.RouteForm) (err error) {
	route := &model.Route{
		Name:      info.Name,
		Type:      info.Type,
		Url:       info.Url,
		UpdatedAt: time.Now().Unix(),
	}
	if info.Id > 0 {
		return rs.repo.Update(info.Id, route)
	}
	route.CreatedAt = time.Now().Unix()
	return rs.repo.Create(route)
}
