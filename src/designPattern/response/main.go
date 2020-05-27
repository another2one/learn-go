package main

import (
	"designPattern/response/utils"
	"fmt"
)

// 拆分程序：拆分为不同小功能，可以自己控制和替换为不同实现
func main() {

	resposeChain := make(map[string]utils.ResponseInter, 1)

	resposeChain["loginCheck"] = utils.LoginCheck{}
	resposeChain["nameLenCheck"] = utils.NameLenCheck{}

	request := utils.Request{
		Name:        "lizhi",
		LoginStatus: false,
	}

	for k, v := range resposeChain {
		fmt.Println(k, " checking ......")
		v.Check(&request)
		fmt.Println(k, " check success !!!")
		fmt.Println()
	}

}
