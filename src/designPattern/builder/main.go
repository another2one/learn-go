package main

import "designPattern/builder/utils"

// 拆分，组合：将大元素拆分为小元素，通过小元素不同组合产生不同大元素
// 将菜单拆为食物，不同食物组合为不同套餐
func main() {
	utils.NonVegPepsiMeal().Check()
	utils.VegCokeMeal().Check()
}
