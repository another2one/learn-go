package common

import(
	"fmt"
)

func PrintStr(str interface{}) {
	fmt.Printf("str's type is %T, str = %q \n", str, str)
}

func PrintOther(other interface{}) {
	fmt.Printf("str's type is %T, str = %v \n", other, other)
}