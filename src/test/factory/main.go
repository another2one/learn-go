package main

import (
	"test/factory/utils"
)

func main() {

	cat := utils.Create("cat")
	cat.Do()

	dog := utils.Create("dog")
	dog.Do()
}
