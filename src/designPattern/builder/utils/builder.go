package utils

import (
	"designPattern/builder/utils/item/burger"
	"designPattern/builder/utils/item/drink"
)

type MealBuilder struct {
}

func VegCokeMeal() *Meal {
	meal := NewMeal()
	meal.AddItem(burger.NewVegBurger())
	meal.AddItem(drink.NewCoke(drink.TypeCOLD))
	return meal
}

func NonVegPepsiMeal() *Meal {
	meal := NewMeal()
	meal.AddItem(burger.NewChickenBurger())
	meal.AddItem(drink.NewPepsi(drink.TypeCOLD))
	return meal
}
