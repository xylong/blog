package internal

import (
	"errors"
)

type Error struct {
	Err  error
	Msg  string
	Code int
}

// PanicIfErr 将错误直接抛出
func PanicIfErr(err error) {
	if err != nil {
		PanicError(SUCCESS, err)
	}
}

// PanicCode 根据错误码抛出
func PanicCode(code int) {
	if code != SUCCESS {
		PanicErrorWithMsg(code, GetMsg(code))
	}
}

// PanicError 抛出错误时，携带状态码
func PanicError(code int, err error) {
	panic(&Error{
		Err:  err,
		Msg:  err.Error(),
		Code: code,
	})
}

// PanicErrorWithMsg 自定义状态码和错误信息的错误
func PanicErrorWithMsg(code int, message string) {
	panic(&Error{
		Err:  errors.New(message),
		Msg:  message,
		Code: code,
	})
}
