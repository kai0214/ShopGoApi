package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type (
	Response struct {
		Code  int         `json:"code"`
		Count int         `json:"count,omitempty"`
		Msg   string      `json:"msg"`
		Data  interface{} `json:"data,omitempty"`
		List  interface{} `json:"list,omitempty" `
	}
)

const (
	SuccessMsg    string = "SUCCESS"
	FailMsg       string = "FAIL"
	ErrFromBottom int    = 90000
	ErrFromValid  int    = 20001
)

func HttpResponse(w *context.Context, err error) {
	msg := SuccessMsg
	code := 0
	if nil != err {
		msg = FailMsg
		code = ErrFromBottom
		baseError, ok := err.(*baseError)
		if ok && baseError != nil {
			code = baseError.Code
			msg = baseError.Message
		} else {
			msg = err.Error()
		}
	}
	hasIndent := beego.AppConfig.String("RunMode") != beego.PROD
	w.Output.JSON(&Response{
		Code: code,
		Msg:  msg,
	}, hasIndent, false)
	return
}
func HttpResponseList(w *context.Context, data interface{}, err error) {
	msg := SuccessMsg
	code := 0
	if nil != err {
		msg = FailMsg
		code = ErrFromBottom
		baseError, ok := err.(*baseError)
		if ok && baseError != nil {
			code = baseError.Code
			msg = baseError.Message
		} else {
			msg = err.Error()
		}
	}
	hasIndent := beego.AppConfig.String("RunMode") != beego.PROD
	w.Output.JSON(&Response{
		Code: code,
		Msg:  msg,
		List: data,
	}, hasIndent, false)
	return
}
func HttpResponseData(w *context.Context, data interface{}, err error) {
	msg := SuccessMsg
	code := 0
	if nil != err {
		if err.Error() == "<QuerySeter> no row found" {
			msg = "数据不存在"
			code = 201
		} else {
			msg = FailMsg
			code = ErrFromBottom
			baseError, ok := err.(*baseError)
			if ok && baseError != nil {
				code = baseError.Code
				msg = baseError.Message
			} else {
				msg = err.Error()
			}
		}
	}
	hasIndent := beego.AppConfig.String("RunMode") != beego.PROD
	w.Output.JSON(&Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}, hasIndent, false)
	return
}
func HttpPageResponse(w *context.Context, count int, data interface{}, err error) {
	msg := SuccessMsg
	code := 0
	if nil != err {
		msg = FailMsg
		code = ErrFromBottom
		baseError, ok := err.(*baseError)
		if ok && baseError != nil {
			code = baseError.Code
			msg = baseError.Message
		} else {
			msg = err.Error()
		}
	}
	hasIndent := beego.AppConfig.String("RunMode") != beego.PROD
	w.Output.JSON(&Response{
		Code:  code,
		Count: count,
		Msg:   msg,
		List:  data,
	}, hasIndent, false)
	return
}
