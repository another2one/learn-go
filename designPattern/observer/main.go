package main

import "learn-go/designPattern/observer/utils"

// 解耦：代码分层，减少依赖
// 消息队列
func main() {
	ob := utils.NewObserver()
	client := &utils.Client{}
	ob.Bind(client)
	s := utils.Subject{Observer: ob}
	s.SetStatus(1)
}
