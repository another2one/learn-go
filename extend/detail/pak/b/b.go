package b

import (
	"fmt"
	"learn-go/extend/detail/pak/c"
)

var A = func() string {
	fmt.Println("b.a ...")
	return c.C
}()

var C = "666"

func init() {
	fmt.Println("b.init ...")
}
