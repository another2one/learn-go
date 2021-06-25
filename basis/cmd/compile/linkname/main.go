package main

import (
	"learn-go/basis/cmd/compile/linkname/priv"
	_ "unsafe"
)

// call private func
//go:linkname private_func learn-go/basis/cmd/compile/linkname/priv.privateFunc
func private_func()

// 调用其他包的公有结构的私有func
//go:linkname public_struc_private_func learn-go/basis/cmd/compile/linkname/priv.(*PublicStruct).f
func public_struc_private_func(p *priv.PublicStruct, a int)

// //go:linkname 绕过编译器检查。直接访问func的实现, priviFunc 和 f 正常时无法访问的
//  警告一下！以下代码均不是常规操作，且存在各种潜在不可控的风险。在项目中应用有可能被同事打死，慎用！！！
func main() {
	private_func() // this is private func
	// 先构造一个other1.PublicStruct
	p := &priv.PublicStruct{I: 1}
	public_struc_private_func(p, 100) // PublicStruct f() 1 0 100
}
