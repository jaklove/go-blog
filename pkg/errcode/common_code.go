package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success = NewError(0,"成功")
	ServerError = NewError(10000000,"服务内部错误")
	InvalidParams = NewError(100000001,"入参错误")
	NotFound = NewError(10000002,"找不到")
	UnauthorizedAuthNotExist = NewError(100000003,"鉴权失败,找不到对应的appKey和appSecret")
	UnauthorizedTokenError = NewError(100000004,"鉴权失败,Token错误")
	UnauthorizedTokenTimeout = NewError(100000005,"鉴权失败，Token失效")
	UnauthorizedGenerate = NewError(10000006,"鉴权失败,Token生成失败")
	TooManyRequests = NewError(10000007,"请求过多")

)

type Error struct {
	code int `json:"code"`
	msg string `json:"msg"`
    details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int,msg string)*Error  {
	if _,ok := codes[code];ok{
		panic(fmt.Sprintf("错误码 %d已经存在，请更换一个",code))
	}
	codes[code] = msg
	return &Error{
		code: code,
		msg: msg,
	}
}

func (e *Error)Error() string  {
	return fmt.Sprintf("错误码：%d,错误信息：%s",e.code,e.msg)
}

func (e *Error)Code() int  {
	return  e.code
}

func (e *Error)Msg() string  {
	return  e.msg
}

func (e *Error)MsgF(args []interface{})string  {
	return fmt.Sprintf(e.msg,args...)
}

func (e *Error)Details()[]string  {
	return  e.details
}

func (e *Error)WithDetails(details ...string)*Error  {
	e.details = []string{}
	for _,d := range details{
		e.details = append(e.details,d)
	}
   return e
}

func (e *Error)StatusCode()int  {
	switch e.code {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return  http.StatusInternalServerError
	case InvalidParams.code:
		return  http.StatusBadRequest
	case UnauthorizedAuthNotExist.code:
		fallthrough
	case UnauthorizedTokenError.code:
		fallthrough
	case UnauthorizedGenerate.code:
		fallthrough
	case UnauthorizedTokenTimeout.code:
		return  http.StatusUnauthorized
	case TooManyRequests.code:
		return  http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}




