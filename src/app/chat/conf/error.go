package conf

import "errors"

var (
	// 用户
	ErrorUserNotExists    = errors.New("user not found")
	ErrorUserExists       = errors.New("user has exists")
	ErrorUserPasswordFail = errors.New("password incorrect")
	ErrorRegisterFail     = errors.New("register fail")
	ErrorUserNotOnline    = errors.New("not online")

	// 通讯
	ErrorMsgType   = errors.New("message type error")
	ErrorMsgLength = errors.New("read content length error")
	ErrorReadError = errors.New("read message error")

	// 系统
	ErrorUnmarshal    = errors.New("unmarshal error")
	ErrorMarshal      = errors.New("marshal error")
	ErrorRedisGetData = errors.New("redis get data error")
	ErrorConnect      = errors.New("connect error")
)

var ErrorCode = map[error]int{

	// 用户
	ErrorUserNotExists:    301,
	ErrorUserExists:       302,
	ErrorUserPasswordFail: 303,
	ErrorRegisterFail:     304,
	ErrorUserNotOnline:    305,

	// 通讯
	ErrorMsgType:   401,
	ErrorMsgLength: 402,
	ErrorReadError: 403,

	// 系统
	ErrorUnmarshal:    501,
	ErrorMarshal:      502,
	ErrorRedisGetData: 503,
	ErrorConnect:      504,
}
