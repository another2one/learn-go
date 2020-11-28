package main

import (
	factory "learn-go/designPattern/factory/utils"
)

// 规范化类：做一个入口来统一产生类
// 所有类必须满足条件，才能被工厂接受
func main() {

	cat := factory.Create("cat")
	cat.Do()

	dog := factory.Create("dog")
	dog.Do()
}
