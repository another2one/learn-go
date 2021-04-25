package c

import (
	"fmt"
)

func init() {
	fmt.Println("c.init ...")
}

var A = func() string {
	fmt.Println("c.a ...")
	return C
}()

var C = "666"
