package service

import (
	"reflect"
	"time"

	"github.com/hechen0210/utils/helper"
)

// 出现key相同的处理方式，cover 覆盖,pre 添加前缀,ignore 忽略 默认 忽略
// pre 添加前缀处理方式时的前缀,新key为 pre_key
type Formater struct {
	Opt    string
	Pre    string
	Result map[string]interface{}
}

const (
	OPTBY_COVER  = "cover"
	OPTBY_PRE    = "pre"
	OPTBY_IGNORE = "ignore"
)

func (p *Formater) Format(data ...interface{}) *Formater {
	result := make(map[string]interface{})
	for _, item := range data {
		_data := helper.Struct2Map(item)
		for key, value := range _data {
			key = helper.HumpToUnder(key)
			if _, exist := result[key]; !exist {
				result[key] = value
			} else {
				switch p.Opt {
				case "cover":
					result[key] = value
				case "pre":
					result[p.Pre+"_"+key] = value
				}
			}
		}
	}
	if value, exist := result["created_at"]; exist && reflect.TypeOf(value).Name() == "int64" {
		result["created_at"] = time.Unix(value.(int64), 0).Format("2006-01-02 15:04:05")
	}
	if value, exist := result["updated_at"]; exist && reflect.TypeOf(value).Name() == "int64" {
		result["updated_at"] = time.Unix(value.(int64), 0).Format("2006-01-02 15:04:05")
	}
	p.Result = result
	return p
}

func (p *Formater) Append(data ...map[string]interface{}) *Formater {
	for _, item := range data {
		for key, value := range item {
			key = helper.HumpToUnder(key)
			p.Result[key] = value
		}
	}
	return p
}
