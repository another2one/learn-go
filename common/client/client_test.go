package client

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	sendData := make(map[string]string)
	sendData["name"] = "测试"
	sendData["sex"] = "男"
	jsonStr, _ := json.Marshal(sendData)

	//此处需要换成你自己的接口地址
	httpUrl := "https://api-local.lewaimai.com"
	headerData := make(map[string]string)
	headerData["X-Ca-Key"] = "22527885"
	headerData["Content-Type"] = "application/json;charset=UTF-8"
	headerData["Accept"] = "application/json"
	body := POST(httpUrl, headerData, jsonStr)
	fmt.Printf("请求成功返回：%s\n", body)
}
