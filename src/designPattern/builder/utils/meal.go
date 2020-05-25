package utils

import (
	"designPattern/builder/utils/item"
	"fmt"
)

type Meal struct {
	Name  string
	Price float64
	items []item.ItemInter
}

func NewMeal() *Meal {
	meal := new(Meal)
	meal.items = make([]item.ItemInter, 0)
	return meal
}

// 价格
func (meal *Meal) CalcPrice() {
	for _, food := range meal.items {
		meal.Price += food.GetPrice()
	}
}

// 打印套餐
func (meal *Meal) Check() {
	fmt.Printf("请核对您的套餐 %s:\n", meal.Name)
	for i, food := range meal.items {
		fmt.Println(i+1, " : ", food.GetName())
	}
	meal.CalcPrice()
	fmt.Printf("总价格 %.2f:\n", meal.Price)
}

// 点餐
func (meal *Meal) AddItem(item item.ItemInter) {
	meal.items = append(meal.items, item)
}
