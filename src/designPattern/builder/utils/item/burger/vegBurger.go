package burger

import "designPattern/builder/utils/item"

type VegBurger struct {
	burger
}

func NewVegBurger() item.ItemInter {
	cb := &VegBurger{}
	cb.Name = "vegBurger"
	cb.Price = 9.8
	return cb
}
