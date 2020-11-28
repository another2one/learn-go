package drink

import (
	"learn-go/designPattern/builder/utils/item"
	"learn-go/designPattern/builder/utils/pack"
)

const (
	TypeNORMAL = iota
	TypeCOLD
	TypeHOT
)

type drink struct {
	item.Item
	Type int
}

func (burger *drink) GetPack() pack.PackingInter {
	return new(pack.Bottle)
}

func (burger *drink) GetDrinkType() string {
	return [...]string{"normal", "cold", "hot"}[burger.Type]
}
