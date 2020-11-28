package main

import "learn-go/designPattern/strategy/utils"

// 一个功能对应多个实现时
// 接口实现
func main() {
	destination := "武汉"
	var trans utils.TravelInter
	trans = &utils.Airport{}
	trans.Travel(destination)
	trans = &utils.Bus{}
	trans.Travel(destination)
}
