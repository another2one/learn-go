package priv

import "fmt"

func privateFunc() {
	fmt.Println("this is private func")
}

type PublicStruct struct {
	I int
	b int
}

func (p *PublicStruct) f(a int) {
	println("PublicStruct f()", p.I, p.b, a)
}
