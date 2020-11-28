package item

import "learn-go/designPattern/builder/utils/pack"

type ItemInter interface {
	GetPack() pack.PackingInter
	GetPrice() float64
	GetName() string
}

type Item struct {
	Name  string
	Price float64
}

func (item Item) GetPrice() float64 {
	return item.Price
}

func (item Item) GetName() string {
	return item.Name
}
