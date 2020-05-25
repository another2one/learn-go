package main

import (
	"designPattern/single/utils"
	"fmt"
)

func main() {

	dog := utils.NewDog()
	dog.Name = "顺子"
	fmt.Println(dog.Name)

	dog2 := utils.NewDog()
	fmt.Println(dog2.Name)

	fmt.Printf("dog = %p, dog2 = %p \n", dog, dog2)
}
