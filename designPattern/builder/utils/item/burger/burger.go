package burger

import (
	"learn-go/designPattern/builder/utils/item"
	"learn-go/designPattern/builder/utils/pack"
)

type burger struct {
	item.Item
}

func (burger *burger) GetPack() pack.PackingInter {
	return new(pack.Wrapper)
}
