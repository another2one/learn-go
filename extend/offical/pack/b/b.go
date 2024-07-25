package b

import (
	"fmt"
	"learn-go/extend/offical/pack/c"
)

var A = func() string {
	fmt.Println("b.a ...")
	return c.C
}()

var C = "666"

func init() {
	fmt.Println("b.init ...")
}
