package service

import (
	"errors"
	"fmt"
	"hr/config"
	"hr/model"
	"hr/repository"
	"hr/request"
	"hr/util"

	"github.com/shopspring/decimal"
)

type DemandService interface {
	GetList(page, perPage int, condition []repository.Condition) (util.Page, error)
	GetInfo(id int) (map[string]interface{}, error)
	Create(data request.DemandForm) (int, error)
	Delete(id int) error
}

type demandServcice struct {
	demandRepo repository.DemandRepository
}

func NewDemandService() DemandService {
	return &demandServcice{
		demandRepo: repository.NewDemandRepository(),
	}
}

// 获取企业需求列表
func (d *demandServcice) GetList(page, perPage int, condition []repository.Condition) (info util.Page, err error) {
	total, err := d.demandRepo.GetCount(condition)
	if err != nil {
		return info, err
	}
	info = util.NewPage(total, perPage, page, []repository.DemandInfo{})
	if total > 0 {
		list, err := d.demandRepo.GetList(page, perPage, condition)
		if err != nil {
			return info, err
		}
		data := []map[string]interface{}{}
		for _, item := range list {
			otherInfo := map[string]interface{}{
				"type_name":   config.WORK_TYPE[item.Type],
				"flag_name":   config.COMPANY_DEMAND_STATUS[item.Flag],
				"expire_name": config.DEMAND_LIMIT["TIME"][item.Expire],
			}
			data = append(data, d.formatData(item, otherInfo))
		}
		info.List = data
	}
	return info, err
}

// 获取企业需求详情
func (d *demandServcice) GetInfo(id int) (info map[string]interface{}, err error) {
	if id == 0 {
		return info, errors.New(config.QUERY_FAIL_MESSAGE)
	}
	result, err := d.demandRepo.GetInfo(id)
	if err != nil {
		return info, err
	}
	data := map[string]interface{}{
		"type":           config.WORK_TYPE[result.Type],
		"flag_name":      config.COMPANY_DEMAND_STATUS[result.Flag],
		"expire":         config.DEMAND_LIMIT["TIME"][result.Expire],
		"published_name": config.DEMAND_PUBLISH_STATUS[result.Published],
	}
	if result.Expire == config.TIME_LIMIT {
		data["expire"] = result.ExpireTime
	}
	return d.formatData(result, data), nil
}

// 创建企业需求
func (d *demandServcice) Create(data request.DemandForm) (id int, err error) {
	base := model.CompanyDemand{
		Name: data.Name,
		Desc: data.Desc,
		Unit: decimal.NewFromFloat32(data.Unit),
		Flag: data.Flag,
	}
	info := model.DemandInfo{
		Type:       data.Type,
		Quantity:   data.Quantity,
		SignTime:   data.SignTime,
		Education:  data.Education,
		Gender:     data.Gender,
		Age:        data.Age,
		Expire:     data.Expire,
		ExpireTime: data.ExpireTime,
	}
	fmt.Println(data)
	if data.Id <= 0 {
		d.demandRepo.Create(base, info)
	} else {
		d.demandRepo.Update(data.Id, base, info)
	}

	return id, err
}

func (d *demandServcice) Delete(id int) error {
	return d.demandRepo.Delete(id)
}

// 格式数据
func (d *demandServcice) formatData(info repository.DemandInfo, more ...map[string]interface{}) (data map[string]interface{}) {
	parse := (&Formater{}).Format(info.CompanyDemand, info.DemandInfo).Append(more...)
	return parse.Result
}
