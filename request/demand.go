package request

import (
	"hr/config"
	"hr/repository"
	"sort"

	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type DemandRequest struct{}

type DemandForm struct {
	Id         int     `validate:"-"`
	Name       string  `validate:"required" message:"请输入需求名称"`
	Desc       string  `validate:"required" message:"请输入需求描述"`
	Type       int     `validate:"required"`
	Unit       float32 `validate:"required|float"`
	Quantity   int     `validate:"required|int|min:1"`
	SignTime   string  `validate:"string" form:"sign_time"`
	Education  int
	Gender     int
	Age        string
	Expire     int    `validate:"required"`
	ExpireTime string `validate:"requiredIf:expire,2" form:"expire_time" message:"请输入需求有效期"`
	Flag       int    `validate:"required|enum:[0,1]"`
}

func NewDemandRequest() *DemandRequest {
	return &DemandRequest{}
}

func (dr *DemandRequest) Search(data map[string]string) []repository.Condition {
	return []repository.Condition{}
}

func (df *DemandRequest) GetFormData(ctx iris.Context) (demandForm DemandForm, err error) {
	err = ctx.ReadForm(&demandForm)
	if err != nil {
		return
	}
	v := validate.Struct(df)
	v.AddRule("Type", "enum", demandForm.typeArgs())
	v.AddRule("Expire", "enum", demandForm.expireArgs())
	return demandForm, v.ValidateE().OneError()
}

func (df *DemandForm) typeArgs() []int {
	data := []int{}
	for index := range config.WORK_TYPE {
		data = append(data, index)
	}
	sort.Ints(data)
	return data
}

func (df *DemandForm) expireArgs() []int {
	data := []int{}
	for index := range config.DEMAND_LIMIT["TIME"] {
		data = append(data, index)
	}
	sort.Ints(data)
	return data
}
