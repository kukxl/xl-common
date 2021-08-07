package http

import (
	"reflect"
	"strings"
	"xl-common/pkg/errs"
)

type Result struct {
	Code 		int
	Message 	string
	Count 		int
	Data 		interface{}
}

func ConvertData(data interface{}) Result {
	return Result{Code: 200, Data: data}
}

func ConvertError(err error) Result {
	typeOf := reflect.TypeOf(err)
	if strings.Contains(typeOf.Name(), "XlErr") {
		e := err.(*errs.XlErr)
		return Result{Code: e.Code, Message: e.Msg}
	}
	return Result{Code: 500, Message: err.Error()}
}

func ConvertPage(count int, data interface{}) Result {
	return Result{Code: 200, Count: count, Data: data}
}