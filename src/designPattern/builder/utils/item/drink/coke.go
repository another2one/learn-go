package drink

import "designPattern/builder/utils/item"

type Coke struct {
	drink
}

func NewCoke(taste int) item.ItemInter {
	cb := &Coke{}
	cb.Type = taste
	cb.Name = cb.GetDrinkType() + " pepsi"
	cb.Price = 2.9
	return cb
}
