package utils

// 类接口
type Class interface {
	Do()
}

var (
	// 保存注册好的工厂信息
	factoryByName = make(map[string]func() Class)
)

func Register(name string, f func() Class) {
	factoryByName[name] = f
}

func Create(name string) Class {
	f, ok := factoryByName[name]
	if !ok {
		panic("not found class")
	}
	return f()
}
