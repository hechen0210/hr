package response

import (
	"hr/config"

	"github.com/kataras/iris/v12"
)

type Response struct {
	Code  int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) SetMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) Fail(ctx iris.Context) {
	if r.Code == 0 {
		r.Code = config.Fail
	}
	if r.Message == "" {
		r.Message = config.Fail_MESSAGE
	}
	ctx.JSON(r)
}

func (r *Response) Success(ctx iris.Context) {
	if r.Code == 0 {
		r.Code = config.SUCCESS
	}
	if r.Message == "" {
		r.Message = config.SUCCESS_MESSAGE
	}
	ctx.JSON(r)
}
