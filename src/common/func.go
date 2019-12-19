package common

import(
	"fmt"
)

func Print(val interface{}) {
	fmt.Printf("val's type is %T, val = %v \n", val, val)
}