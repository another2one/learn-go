package main

import (
	factory "designPattern/factory/utils"
)

func main() {

	cat := factory.Create("cat")
	cat.Do()

	dog := factory.Create("dog")
	dog.Do()
}
