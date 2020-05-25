package main

import "designPattern/builder/utils"

func main() {
	utils.NonVegPepsiMeal().Check()
	utils.VegCokeMeal().Check()
}
