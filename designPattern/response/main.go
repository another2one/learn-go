package main

import (
	"learn-go/designPattern/response/utils"
	"fmt"
)

// 拆分程序：拆分为不同小功能，可以自己控制和替换为不同实现
func main() {

	responseChain := make(map[string]utils.ResponseInter, 1)

	responseChain["loginCheck"] = utils.LoginCheck{}
	responseChain["nameLenCheck"] = utils.NameLenCheck{}

	request := utils.Request{
		Name:        "lizhi",
		LoginStatus: false,
	}

	for k, v := range responseChain {
		fmt.Println(k, " checking ......")
		v.Check(&request)
		fmt.Println(k, " check success !!!")
		fmt.Println()
	}

}
