package utils

import (
	"app/chat/conf"
	"encoding/json"
	"fmt"
)

func GetServerPushData(data string, msgType string) (msgDataBS []byte, err error) {

	// 获取结果数据
	msgData := conf.Msg{
		Type: msgType,
		Data: data,
	}
	msgDataBS, err = json.Marshal(msgData)
	if err != nil {
		err = conf.ErrorMarshal
		return
	}
	return
}

func GetSendData(errorCode error, data string) (bs []byte, err error) {

	code, msg := getCodeAndMsgByError(errorCode)

	resData := conf.Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	resDataBS, err := json.Marshal(resData)
	if err != nil {
		err = conf.ErrorUnmarshal
		return
	}
	msgData := conf.Msg{
		Type: conf.ResultType,
		Data: string(resDataBS),
	}
	bs, err = json.Marshal(msgData)
	if err != nil {
		err = conf.ErrorMarshal
		return
	}
	return
}

func getCodeAndMsgByError(err error) (int, string) {
	if err == nil {
		return 200, "success"
	}
	if _, ok := conf.ErrorCode[err]; ok {
		return conf.ErrorCode[err], err.Error()
	} else {
		fmt.Println("server error: " + err.Error())
		return 500, "server error"
	}
}
