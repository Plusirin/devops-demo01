package myerr

import "fmt"

var (
	// Common errors
	OK                  = &ErrNum{Code: 0, Message: "OK"}
	InternalServerError = &ErrNum{Code: 30001, Message: "内部错误."}
	ErrBind             = &ErrNum{Code: 30002, Message: "请求信息无法转换成结构体."}
	ErrDatabase         = &ErrNum{Code: 30002, Message: "Database error."}
	ErrValidation       = &ErrNum{Code: 30001, Message: "Validation failed."}
	ErrEncrypt          = &ErrNum{Code: 30101, Message: "Error occurred while encrypting the user password."}
	// user errors
	ErrAccountNotFound = &ErrNum{Code: 50102, Message: "用户不存在."}
	ErrPassword        = &ErrNum{Code: 50103, Message: "密码错误."}
	ErrAccountEmpty    = &ErrNum{Code: 50104, Message: "用户名不能为空."}
	ErrPasswordEmpty   = &ErrNum{Code: 50103, Message: "密码不能为空."}
	ErrMissingHeader   = &ErrNum{Code: 50104, Message: "Header 不存在"}
	ErrToken           = &ErrNum{Code: 50105, Message: "生成 Token 错误"}
	PassParamCheck     = &ErrNum{Code: 60000, Message: "参数校验通过"}
)

type ErrNum struct {
	Code    int
	Message string
}

func (e *ErrNum) Error() string {
	return e.Message
}

type Err struct {
	ErrNum ErrNum
	Err    error
}

func New(num ErrNum, err error) *Err {
	return &Err{
		ErrNum: ErrNum{Code: num.Code, Message: num.Message},
		Err:    err,
	}
}

func (e *Err) Add(message string) Err {
	e.ErrNum.Message += " " + message
	return *e
}

func (err *Err) AddFormat(format string, args ...interface{}) Err {
	err.ErrNum.Message += " " + fmt.Sprintf(format, args...)
	return *err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.ErrNum.Code, err.ErrNum.Message, err.Err)
}

func IsErrAccountNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrAccountNotFound.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.ErrNum.Code, typed.ErrNum.Message
	case *ErrNum:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
