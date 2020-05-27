package main

import "designPattern/decorator/utils"

func main() {
	var decorator utils.DecoratorInter
	decorator = &utils.Decorator{}
	decorator.Cache(new(utils.Line))
}
