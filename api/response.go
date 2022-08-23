package api

import (
	"github.com/kataras/iris/v12"
)

type Response struct {
	Ctx     iris.Context `json:"-"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

func NewResponse(ctx iris.Context) Response {
	return Response{Ctx: ctx}
}

/*
Success 返回成功
message 信息
data 返回的数据
*/
func (r *Response) Success(message string) {
	r.Code = 200;
	r.Message = message
	r.Ctx.JSON(r)
}

func (r *Response) Fail(code int, message string) {
	r.Code = code
	r.Message = message
	r.Ctx.JSON(r)
}

func (r *Response) SetData(obj interface{}, need []string) *Response {
	// t := reflect.TypeOf(obj)
	// v := reflect.ValueOf(obj)
	// var data = make(map[string]interface{})
	// for i := 0; i < t.NumField(); i++ {
	// 	if v.Field(i).Interface() != 0 && v.Field(i).Interface() != "" {
	// 		key := t.Field(i).Name
	// 		if len(need) == 0 || helper.Contains(need, key) {
	// 			data[key] = v.Field(i).Interface()
	// 		}
	// 	}
	// }
	r.Data = obj
	return r
}
