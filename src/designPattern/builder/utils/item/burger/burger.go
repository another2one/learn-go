package burger

import (
	"designPattern/builder/utils/item"
	"designPattern/builder/utils/pack"
)

type burger struct {
	item.Item
}

func (burger *burger) GetPack() pack.PackingInter {
	return new(pack.Wrapper)
}
