package drink

import "designPattern/builder/utils/item"

type Pepsi struct {
	drink
}

func NewPepsi(taste int) item.ItemInter {
	cb := &Pepsi{}
	cb.Type = taste
	cb.Name = cb.GetDrinkType() + " pepsi"
	cb.Price = 2.9
	return cb
}
