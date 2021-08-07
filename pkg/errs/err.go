package errs

import "fmt"

type XlErr struct {
	Code int
	Msg string
}

var (
	ErrSystem 		= newErr(1000, "系统异常")
	ErrPwdNotMath 	= newErr(2000, "密码错误")
)

func (e *XlErr) Error() string {
	return fmt.Sprintf("xlErr(code: %d, msg: %s)", e.Code, e.Msg)
}

func newErr(code int, msg string) *XlErr {
	return &XlErr{Code: code, Msg: msg}
}