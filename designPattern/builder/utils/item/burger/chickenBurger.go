package burger

import "learn-go/designPattern/builder/utils/item"

type ChickenBurger struct {
	burger
}

func NewChickenBurger() item.ItemInter {
	cb := &ChickenBurger{}
	cb.Name = "chickenBurger"
	cb.Price = 13.8
	return cb
}
